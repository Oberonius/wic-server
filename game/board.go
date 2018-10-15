package game

import (
	"errors"
	"fmt"
	"sync"
	"wic-server/interfaces"
)

const (
	maxX         = 10
	maxY         = 30
	archerJoined = "JOINED"
	archerLeaved = "LEAVED"
)

type Board struct {
	sync.RWMutex

	gameName     string
	actors       map[string]interfaces.Actor
	commChannels map[string]interfaces.Communicator
	inGame       bool

	stateChangeObserver func(*Board)
}

func NewBoard(gameName string, observer func(*Board)) *Board {
	return &Board{
		gameName:            gameName,
		actors:              make(map[string]interfaces.Actor),
		commChannels:        make(map[string]interfaces.Communicator),
		inGame:              false,
		stateChangeObserver: observer,
	}
}

func (b *Board) IsInGame() bool {
	b.RLock()
	defer b.RUnlock()
	return b.inGame
}

func (b *Board) GetGameName() string {
	b.RLock()
	defer b.RUnlock()
	return b.gameName
}

func (b *Board) NotifyLeaved(p interfaces.Actor) {
	b.Lock()
	if b.inGame && p.GetRole() == interfaces.ActorRoleArcher {
		b.sendArcherState(archerLeaved, p)
	}
	delete(b.actors, p.GetName())
	delete(b.commChannels, p.GetName())

	if len(b.actors) == 0 {
		b.Unlock()
		b.Stop()
	} else {
		b.Unlock()
	}
}

//AddPlayer uses user name as a unique identifier inside the game
//it checks uniqueness of the username and adds new player
//all added players will receive information about current game events
func (b *Board) AddPlayer(p interfaces.Actor, comm interfaces.Communicator) error {
	b.Lock()
	defer b.Unlock()

	name := p.GetName()
	if _, ok := b.actors[name]; ok {
		return errors.New(fmt.Sprintf("player '%s' already exists in this game", name))
	}
	b.actors[name] = p
	b.commChannels[name] = comm

	if b.inGame {
		p.Spawn()
		b.sendArcherState(archerJoined, p)
	}

	return nil
}

func (b *Board) sendArcherState(state string, actor interfaces.Actor) {
	b.notifyExcluding(func(c interfaces.Communicator) {
		c.NotifyArcherState(state, actor)
	}, actor.GetName())
}

func (b *Board) Run() {
	if b.inGame {
		return
	}

	for _, a := range b.actors {
		a.Spawn()
	}

	b.inGame = true
	b.stateChangeObserver(b)
}

func (b *Board) Stop() {
	b.Lock()
	for _, a := range b.actors {
		a.Stop()
	}
	b.inGame = false
	observer := b.stateChangeObserver
	b.Unlock()

	observer(b)
}

func (b *Board) GetMaxX() int {
	return maxX
}

func (b *Board) GetMaxY() int {
	return maxY
}
