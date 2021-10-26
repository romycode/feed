package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
)

type TCPServer struct {
	addr    string
	maxConn int
}

func NewTCPServer(addr string, maxConn int) *TCPServer {
	return &TCPServer{addr: addr, maxConn: maxConn}
}

func (s *TCPServer) Start(ctx context.Context) error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("cannot start tcp listener in addr: %s, err: %w", s.addr, err)
	}

	for i := 0; i < s.maxConn; i++ {
		go handle(ctx, ln)
	}

	select {
	case <-ctx.Done():
		return nil
	}
}

func handle(ctx context.Context, ln net.Listener) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			conn, err := ln.Accept()
			if err != nil {
				log.Println("Error accepting request on port 4000: " + err.Error())
			}

			data, err := io.ReadAll(conn)
			if err != nil {
				log.Println("error reading request on port 4000: " + err.Error())
			}
			rawSKU := string(data)

			fmt.Println("SKU: " + rawSKU)

			_ = conn.Close()
		}
	}
}
