package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/nikkomiu/forza-utel/server"
)

type Server struct {
	*server.Server

	w io.WriteCloser
}

func NewServer() *Server {
	f, _ := os.OpenFile("data.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	return &Server{
		Server: server.NewDefaultServer(),
		w:      f,
	}
}

func (s Server) handleData(addr net.Addr, data interface{}) error {
	return json.NewEncoder(s.w).Encode(data)
}

func (s Server) handleError(addr net.Addr, data interface{}, err error) {
	fmt.Printf("error handling data from %s: %s\n", addr, err)
}

func (s *Server) Listen(address string) error {
	s.Handler = s.handleData
	s.ErrorHandler = s.handleError

	log.Println("Starting server...")
	return s.Server.Listen(address)
}

func main() {
	if err := NewServer().Listen(":4040"); err != nil {
		fmt.Fprintf(os.Stderr, "error running server: %s\n", err)
		os.Exit(1)
	}
}
