package xlt

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type Request struct {
	Method     string    // 请求方法
	Url        *url.URL  // url
	Proto      string    // 协议
	Header     Header    // 请求头
	Body       io.Reader // 请求体
	RemoteAddr string    // 客户端地址
	RequestURI string    // 字符串形式的url

	conn        *conn             // 产生此Request的Http连接
	cookies     map[string]string // 存储cookie
	queryString map[string]string // 存储查询键值
}

func readRequest(c *conn) (*Request, error) {
	r := &Request{}
	r.conn = c
	r.RemoteAddr = r.conn.rwc.RemoteAddr().String()

	// read first line, example: Get /index?name=xlt HTTP/1.1
	line, err := readLine(c.bufr)
	if err != nil {
		return r, err
	}

	log.Printf("line data: [%s]\n", string(line))
	n, err := fmt.Sscanf(string(line), "%s%s%s", &r.Method, &r.RequestURI, &r.Proto)
	if err != nil {
		return r, err
	}
	log.Printf("n: [%d], method: [%s],  requesturi: [%s], proto: [%s]\n", n, r.Method, r.RequestURI, r.Proto)

	r.Url, err = url.ParseRequestURI(r.RequestURI)
	if err != nil {
		return r, err
	}
	log.Printf("parse_request_url:[%+v]\n", r.Url)

	r.parseQuery()
	log.Printf("parse_query:[%+v]\n", r.queryString)

	r.Header, err = readHeader(c.bufr)
	if err != nil {
		return r, err
	}

	const noLimit = (1 << 63) - 1
	r.conn.lr.N = noLimit
	r.setupBody()

	return r, nil
}

//parseQuery example: 127.0.0.1?name=xlt&token=12345
func (r *Request) parseQuery() {
	r.queryString = parseQuery(r.Url.RawQuery)
}

func (r *Request) chunked() bool {
	te := r.Header.Get("Transfer-Encoding")
	return te == "chunked"
}

func (r *Request) setupBody() {
	//只允许POST和PUT方法设置报文主体
	if r.Method != "POST" && r.Method != "PUT" {
		r.Body = &eofReader{}
	} else if cl := r.Header.Get("Content-Length"); cl != "" {
		// 设置了Content-Length
		contentLength, err := strconv.ParseInt(cl, 10, 64)
		if err != nil {
			r.Body = &eofReader{}
			return
		}
		r.Body = io.LimitReader(r.conn.bufr, contentLength)
		r.fixExpectContinueReader()
	} else {
		r.Body = &eofReader{}
	}
}

func (r *Request) finishRequest() error {
	//将缓存中剩余的数据发送到rwc中
	if err := r.conn.bufw.Flush(); err != nil {
		return err
	}

	// 消费剩余数据
	_, err := io.Copy(ioutil.Discard, r.Body)
	return err
}

func (r *Request) Query(name string) string {
	if r.cookies == nil {
		r.parseCookies()
	}

	return r.cookies[name]
}

//parseCookies 解析请求头中的Cookie
func (r *Request) parseCookies() {
	if r.cookies != nil {
		return
	}

	cookies, ok := r.Header["Cookie"]
	if !ok {
		return
	}

	r.cookies = make(map[string]string, len(cookies))

	for _, value := range cookies {
		// example: uuid=123456789; tid=a1b2c3d4e5f6g7; HOME=1
		kvs := strings.Split(strings.TrimSpace(value), ";")
		if len(kvs) == 1 && kvs[0] == "" {
			continue
		}

		for i := 0; i < len(kvs); i++ {
			// example: uuid=123456789
			index := strings.Index(kvs[i], "=")
			if index == -1 {
				continue
			}

			r.cookies[strings.TrimSpace(kvs[i][:index])] = strings.TrimSpace(kvs[i][index+1:])
		}
	}

	return
}

//readLine 读取一行数据
//isPrefix 如果为true，代表请求头还没读完需要继续
func readLine(bufr *bufio.Reader) ([]byte, error) {
	data, isPrefix, err := bufr.ReadLine()
	if err != nil {
		return data, err
	}

	over := make([]byte, 0, 10)
	for isPrefix {
		over, isPrefix, err = bufr.ReadLine()
		if err != nil {
			break
		}
		data = append(data, over...)
	}

	return data, err
}

//RawQuery example: name=xlt&token=12345
func parseQuery(RawQuery string) map[string]string {
	// 以 & 符号分隔得到查询键值对
	parts := strings.Split(RawQuery, "&")
	queries := make(map[string]string, len(parts))

	for _, part := range parts {
		// 以 = 符号分割得到键值存储到map
		index := strings.Index(part, "=")
		if index == -1 || index == len(parts)-1 {
			continue
		}
		queries[strings.TrimSpace(part[:index])] = strings.TrimSpace(part[index+1:])
	}

	return queries
}

//readHeader 读取请求头数据
func readHeader(bufr *bufio.Reader) (Header, error) {
	header := Header{}

	for {
		line, err := readLine(bufr)
		// 没有读取到任何数据，会返回err
		if err != nil {
			return nil, err
		}

		// 读取到0字节数据，请求头结束
		if len(line) == 0 {
			break
		}

		//以 ":" 符号分割键值对存储
		index := bytes.IndexByte(line, ':')
		if index == -1 {
			return nil, errors.New("unsupported protocol")
		}
		if index == len(line)-1 {
			continue
		}

		key, value := string(line[:index]), strings.TrimSpace(string(line[index+1:]))
		header[key] = append(header[key], value)
	}

	return header, nil
}

type eofReader struct {
}

func (er *eofReader) Read([]byte) (n int, err error) {
	return 0, io.EOF
}

type expectContinueReader struct {
	wroteContinue bool
	r io.Reader
	w *bufio.Writer
}

func (er *expectContinueReader) Read(p []byte) (n int, err error) {
	//第一次读取前发送100 continue
	if !er.wroteContinue {
		er.w.WriteString("HTTP/1.1 100 Continue\r\n\r\n")
		er.w.Flush()
		er.wroteContinue = true
	}

	return er.r.Read(p)
}

func (r *Request) fixExpectContinueReader() {
	if r.Header.Get("Expect") != "100-continue" {
		return
	}

	r.Body = &expectContinueReader{
		r: r.Body,
		w: r.conn.bufw,
	}
}

