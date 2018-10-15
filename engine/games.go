package engine

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"wic-server/game"
	"wic-server/interfaces"
	"wic-server/players"
)

type Games struct {
	sync.Mutex
	games map[string]*game.Board
}

func NewGamesPool() *Games {
	return &Games{
		games: map[string]*game.Board{},
	}
}

func (g *Games) StartGame(playerName string, comm interfaces.Communicator) (interfaces.Actor, error) {
	g.Lock()
	defer g.Unlock()

	if _, ok := g.games[playerName]; ok {
		return nil, errors.New(fmt.Sprintf("game '%s' already exists", playerName))
	}

	board := game.NewBoard(playerName, g.GameStateChangeObserver)
	g.games[playerName] = board

	zombie := players.NewZombie(rand.Intn(board.GetMaxX()+1), 0, "DEAD-KNIGHT", board)
	archer := players.NewArcher(playerName, board)

	board.AddPlayer(zombie, nil)
	board.AddPlayer(archer, comm)
	board.Run()

	return archer, nil
}

func (g *Games) JoinGame(gameName, playerName string, comm interfaces.Communicator) (interfaces.Actor, error) {
	g.Lock()
	defer g.Unlock()

	s, ok := g.games[gameName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("game '%s' does not exist", gameName))
	}

	archer := players.NewArcher(playerName, s)
	if err := s.AddPlayer(archer, comm); err != nil {
		return nil, err
	}

	return archer, nil
}

func (g *Games) GameStateChangeObserver(board *game.Board) {
	if !board.IsInGame() {
		g.Lock()
		defer g.Unlock()
		delete(g.games, board.GetGameName())
	}
}
