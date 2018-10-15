package tcpServer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
	"wic-server/engine"
	"wic-server/interfaces"
)

const (
	cmdStart = "start"
	cmdJoin  = "join"
	cmdShoot = "shoot"
	cmdExit  = "exit"

	ErrNoCommand       = "no command"
	ErrWrongParameters = "wrong parameters"
	ErrWrongCommand    = "wrong command"
	ErrAlreadyInGame   = "you're already in the game. send 'exit' to leave"
	ErrNotInGame       = "you're not in the game. type 'start {name}' or 'join {game} {name}' to play"
)

type Session struct {
	conn        net.Conn
	reader      *bufio.Reader
	writer      *bufio.Writer
	killSession chan bool
	player      interfaces.Actor
	games       *engine.Games
}

type SessionsManager interface {
	Finished(*Session)
}

func NewSession(conn net.Conn, games *engine.Games) *Session {
	s := &Session{
		conn:        conn,
		games:       games,
		reader:      bufio.NewReader(conn),
		writer:      bufio.NewWriter(conn),
		killSession: make(chan bool, 1),
	}
	return s
}

func (s *Session) Handle() {
	s.writer.Flush()
	for {
		select {
		case <-s.killSession:
			return
		default:
			line, err := s.reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					s.kill()
				}
				time.Sleep(100 * time.Millisecond)
				continue
			}
			if err := s.ProcessCommand(strings.Fields(line)); err != nil {
				fmt.Fprintln(s.writer, "ERROR "+err.Error())
				s.writer.Flush()
			}

		}
	}
}

func (s *Session) Write(p []byte) (n int, err error) {
	n, err = s.writer.Write(p)
	s.writer.Flush()
	return
}

func (s *Session) ProcessCommand(params []string) error {
	var err error

	if len(params) == 0 {
		return errors.New(ErrNoCommand)
	}

	if s.player != nil && s.player.GetState() == interfaces.ActorStateLeaved {
		s.player = nil
	}

	switch strings.ToLower(params[0]) {
	case cmdStart:
		if len(params) < 2 {
			return errors.New(ErrWrongParameters)
		}
		if s.player != nil {
			return errors.New(ErrAlreadyInGame)
		}
		s.player, err = s.games.StartGame(params[1], NewCommunicator(s))
		return err

	case cmdJoin:
		if len(params) < 3 {
			return errors.New(ErrWrongParameters)
		}
		if s.player != nil {
			return errors.New(ErrAlreadyInGame)
		}
		s.player, err = s.games.JoinGame(params[1], params[2], NewCommunicator(s))
		return err

	case cmdShoot:
		if s.player == nil {
			return errors.New(ErrNotInGame)
		}

		if len(params) < 3 {
			return errors.New(ErrWrongParameters)
		}
		x, errX := strconv.Atoi(params[1])
		y, errY := strconv.Atoi(params[2])
		if errX != nil || errY != nil {
			return errors.New(ErrWrongParameters)
		}

		s.player.Shoot(x, y)
		return nil

	case cmdExit:
		if s.player == nil {
			return errors.New(ErrNotInGame)
		}
		s.player.Leave()
		s.player = nil
		return nil
	}

	return errors.New(ErrWrongCommand)
}

func (s *Session) kill() {
	if s.player != nil {
		s.player.Leave()
	}
	s.killSession <- true
}
