package xlt

import (
	"net"
)

type Handler interface {
	ServerHttp(w ResponseWriter, r *Request)
}

type Server struct {
	Addr    string
	Handler Handler
}

func (s *Server) ListenAndServe() error {
	listen, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	for true {
		rwc, err := listen.Accept()
		if err != nil {
			return err
		}

		conn := newConn(rwc, s)
		go conn.serve()
	}

	return nil
}
