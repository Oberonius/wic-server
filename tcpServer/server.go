package tcpServer

import (
	"errors"
	"fmt"
	"net"
	"time"
	"wic-server/engine"
)

type Server struct {
	games *engine.Games
}

func Run(port string, games *engine.Games) error {
	s := &Server{games: games}
	return s.Run(port)
}

func (s *Server) Run(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return errors.New(fmt.Sprintf("Error listening: %v", err.Error()))
	}
	defer listener.Close()

	fmt.Println("Listening on port:", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection accepting failed.")
			conn.Close()
			time.Sleep(100 * time.Millisecond)
			continue
		}

		fmt.Println("A new connection accepted.")
		go NewSession(conn, s.games).Handle()
	}
}
