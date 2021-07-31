package xlt

import (
	"bufio"
	"errors"
	"io"
)

type chunkReader struct {
	//n 当前处理的块中还剩下多少字节
	n               int
	bufr            *bufio.Reader
	//done 记录报文主题是否读取完毕
	done            bool
	crlf            [2]byte
	haveDiscardCRLF bool
}

func (cw *chunkReader) Read(p []byte) (n int, err error) {
	if cw.done {
		return 0, io.EOF
	}

	var nn int
	lenP := len(p)

	//已读取到的数据量小于总需要读取的数据量
	for n < lenP {
		//当前块可读取的数据量大于等于需要读取的数据量
		if len(p) <= cw.n {
			nn, err = cw.bufr.Read(p)
			cw.n -= nn
			return nn, err
		}

		//当前块不能够装下需要读取的数据量
		_, err := io.ReadFull(cw.bufr, p[:cw.n])
		if err != nil {
			return n, err
		}
		n += cw.n
		p = p[cw.n:]

		//需要把当前块的\r\n消费掉
		if err := cw.discardCRLF(); err != nil {
			return n, err
		}

		//获取下一个chunk data的长度
		cw.n, err = cw.getChunkSize()
		if err != nil {
			return n, err
		}
	}
	return
}

//discardCRLF 每一次读取完chunk data之后需要舍弃紧跟的\r\n两个字节
func (cw *chunkReader) discardCRLF() error {
	if !cw.haveDiscardCRLF {
		cw.haveDiscardCRLF = true
		return nil
	}

	// 读取\r\n
	if _, err := io.ReadFull(cw.bufr, cw.crlf[:]); err != nil {
		return errors.New("discard crlf failed, err: " + err.Error())
	}

	if cw.crlf[0] != '\r' || cw.crlf[1] != '\n' {
		return errors.New("unsupported encoding format of chunk")
	}

	return nil
}

func (cw *chunkReader) getChunkSize() (chunkSize int, err error) {
	line, err := readLine(cw.bufr)
	if err != nil {
		return
	}

	// 16进制转换为10进制
	for i := 0; i < len(line); i++ {
		switch {
		case 'a' <= line[i] && line[i] <= 'f':
			chunkSize = chunkSize * 16 + int(line[i] - 'a') + 10
		case 'A' <= line[i] && line[i] <= 'F':
			chunkSize = chunkSize * 16 + int(line[i] - 'A') + 10
		case '0' <= line[i] && line[i] <= '9':
			chunkSize = chunkSize * 16 + int(line[i] - '0')
		default:
			return 0, errors.New("illegal hex number")
		}
	}

	return
}

