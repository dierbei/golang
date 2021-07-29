package xlt

import (
	"bufio"
	"io"
	"log"
	"net"
	"time"
)

// 对于HTTP协议来说，一个请求报文分为三部分：请求行、首部字段以及报文主体，一个post请求的报文如下：
// POST / HTTP/1.1\r\n						   	# 请求行
// Content-Type: text/plain\r\n				    # 2~7 请求头, 键值
// User-Agent: PostmanRuntime/7.28.0\r\n
// Host: 127.0.0.1:8080\r\n
// Accept-Encoding: gzip, deflate, br\r\n
// Connection: keep-alive\r\n
// Content-Length: 18\r\n
// \r\n
// hello,I am client!							# 请求体

// 其中首部字段部分是由一个个key-value对组成，每一对之间通过\r\n分割
// 首部字段与报文主体之间则是利用空行(CR+LF)即\r\n\r\n作为分界
// 首部字段到底有多少个key-value对于服务端程序来说是无法预知的
// 因此我们想正确解析出所有的首部字段，我们必须一直解析到出现CR+LF为止

type conn struct {
	srv *Server

	// conn每次接收到数据都会写入数据，进行一次系统调用的IO操作
	rwc net.Conn

	// 由于conn进行系统调用，对应用程序的性能影响较大，因此引入缓存机制
	bufw *bufio.Writer

	// 对于一个正常的http请求报文，其首部字段总长度不会超过1MB
	// 所以直接不加限制的读到CRLF完全可行，但问题是无法保证所有的客户端都没有恶意
	// LimitedReader底层为Reader
	lr *io.LimitedReader

	// io.LimitedReader封装成bufio.Reader可以使用ReadLine方法
	bufr *bufio.Reader
}

func newConn(rwc net.Conn, srv *Server) *conn {
	// 用户可能在阅读你框架源码后发现你对首部字段的读取未采取任何限制措施
	// 于是发送了一个首部字段无限长的http请求，导致你的电脑无限解析最终用掉了所有内存直至程序崩溃
	// 因此我们应该为我们的reader限制最大读取量
	// 如果总共读取的字节数超过了这个N，则接下来对这个reader的读取都会返回io.EOF错误
	lr := &io.LimitedReader{
		R: rwc,
		N: 1 << 20, // 1MB
	}

	return &conn{
		srv:  srv,
		rwc:  rwc,
		bufw: bufio.NewWriterSize(rwc, 4<<10), // 4k
		lr:   lr,
		bufr: bufio.NewReaderSize(lr, 4<<10),
	}
}

func (c *conn) serve() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic recovered, err:%v\n", err)
		}
		c.Close()
	}()

	// http1.1 支持keep-alive长连接，一个连接可能读取多个请求，因此使用循环读取
	for true {
		req, err := c.readRequest()
		if err != nil {
			log.Printf("[readRequest] failed. err:%v\n", err)
			break
		}

		res := c.setupResponse()
		c.srv.Handler.ServerHttp(res, req)
		time.Sleep(time.Second)

		// 请求数据都写入bufw，缓存默认大小为4k
		// 同时，在一个请求之后，bufw可能还缓存有部分数据，需要调用Flush保证数据全部发送
		if err := c.bufw.Flush(); err != nil {
			log.Printf("flush data failed. err:%v\n", err)
		}
	}
}

// Close close tcp connect
func (c *conn) Close() {
	if err := c.rwc.Close(); err != nil {
		log.Printf("Close conn failed, err:%v\n", err)
	}
}

// readRequest according conn get request
func (c *conn) readRequest() (*Request, error) {
	return readRequest(c)
}

func (c *conn) setupResponse() *response {
	return setupResponse(c)
}
