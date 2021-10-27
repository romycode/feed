package server

import (
	"context"
	"fmt"
	"io"
	"net"
)

type TCPServer struct {
	addr    string
	maxConn int

	output chan string
}

func NewTCPServer(addr string, maxConn int, output chan string) *TCPServer {
	return &TCPServer{addr: addr, maxConn: maxConn, output: output}
}

func (s *TCPServer) Start(ctx context.Context) error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("cannot start tcp listener in addr: %s, err: %w", s.addr, err)
	}
	defer ln.Close()

	errs := make(chan error)

	for i := 0; i < s.maxConn; i++ {
		go handle(ctx, ln, s.output, errs)
	}

	select {
	case err := <-errs:
		return err
	case <-ctx.Done():
		return nil
	}
}

func handle(ctx context.Context, ln net.Listener, output chan string, errs chan error) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			conn, err := ln.Accept()
			if err != nil {
				errs <- fmt.Errorf("error accepting request on port 4000: %w", err)
			}

			data, err := io.ReadAll(conn)
			if err != nil {
				errs <- fmt.Errorf("error reading connection data on port 4000: %w", err)
			}

			output <- string(data)

			_ = conn.Close()
		}
	}
}
