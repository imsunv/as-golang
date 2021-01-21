package function

import (
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

type Option func(*Server)

func Protocol(p string) Option {
	return func(server *Server) {
		server.Protocol = p
	}
}

func Timeout(t time.Duration) Option {
	return func(server *Server) {
		server.Timeout = t
	}
}

func NewServer(addr string, port int, options ...Option) (*Server, error) {
	server := Server{
		Addr: addr,
		Port: port,
	}

	for _, option := range options {
		option(&server)
	}

	return &server, nil
}
