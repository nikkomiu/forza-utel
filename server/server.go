package server

import (
	"fmt"
	"log"
	"net"

	"github.com/nikkomiu/forza-utel/format"
)

var ErrorNoMatchingFormat = fmt.Errorf("no matching format found")

type MessageHandleFunc func(addr net.Addr, data interface{}) error
type ErrorMessageHandleFunc func(addr net.Addr, data interface{}, err error)

type FormatSerializer interface {
	// Parse should parse the data into this format. If parsing the message fails the server
	// will continue to try other parsers where CanParse return true
	Parse(data []byte, len int) (interface{}, error)
}

type Server struct {
	udpServer net.PacketConn

	formats []FormatSerializer

	Handler      MessageHandleFunc
	ErrorHandler ErrorMessageHandleFunc
}

func NewDefaultServer() *Server {
	return NewServer(
		format.NewDashSerializer(),
		format.NewSLEDSerializer(),
	)
}

func NewServer(formats ...FormatSerializer) *Server {
	return &Server{formats: formats}
}

func (s *Server) RegisterFormat(format ...FormatSerializer) {
	s.formats = append(s.formats, format...)
}

func (s *Server) Listen(address string) (err error) {
	server, err := net.ListenPacket("udp", address)
	if err != nil {
		return
	}

	s.udpServer = server

	buf := make([]byte, 512)
	for {
		n, addr, err := s.udpServer.ReadFrom(buf)
		if err != nil && err.(*net.OpError).Op == "read" {
			return nil
		} else if err != nil {
			log.Printf("error reading data: %#v\n", err)
			continue
		}

		go s.handleMessage(addr, buf[:], n)
	}
}

func (s *Server) Shutdown() error {
	return s.udpServer.Close()
}

func (s Server) handleMessage(addr net.Addr, data []byte, n int) {
	for i, dataFormat := range s.formats {
		data, err := dataFormat.Parse(data, n)
		if err == format.ErrorInvalidLength {
			continue
		} else if err != nil {
			log.Printf("error parsing data: %s\n", err)
			continue
		}

		if i != 0 {
			s.moveFormatToFront(i)
		}

		if err = s.Handler(addr, data); err != nil && s.ErrorHandler != nil {
			s.ErrorHandler(addr, data, err)
		}
	}
}

// moveFormatToFront to move this format to the front of the formats slice.
// This is useful for formats that are more likely to be used.
func (s *Server) moveFormatToFront(idx int) {
	if idx == 0 {
		return
	}
	s.formats[idx], s.formats[0] = s.formats[0], s.formats[idx]
}
