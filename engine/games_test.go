package engine

import (
	"testing"
	"wic-server/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestItStartsNewGame(t *testing.T) {
	pool := NewGamesPool()

	comm := &interfaces.MockCommunicator{}

	arc, err := pool.StartGame("foo", comm)
	assert.Equal(t, 1, len(pool.games))
	assert.Nil(t, err)

	comm.On("NotifyShot", arc, nil).Once()
	arc.Shoot(5, 5)
}

func TestItDeniesDuplicateNames(t *testing.T) {
	pool := NewGamesPool()

	comm := &interfaces.MockCommunicator{}
	_, err := pool.StartGame("foo", comm)

	comm2 := &TestCommunicator{}
	_, err = pool.JoinGame("foo", "foo", comm2)
	assert.NotNil(t, err)

	comm3 := &interfaces.MockCommunicator{}
	_, err = pool.StartGame("foo", comm3)
	assert.NotNil(t, err)
}

func TestItDeniesJoinOnNonExistingGame(t *testing.T) {
	pool := NewGamesPool()

	comm := &interfaces.MockCommunicator{}
	_, err := pool.StartGame("foo", comm)

	comm2 := &TestCommunicator{}
	_, err = pool.JoinGame("bar", "baz", comm2)
	assert.NotNil(t, err)
}

func TestItLetsToJoinGame(t *testing.T) {
	pool := NewGamesPool()

	comm := &TestCommunicator{}

	_, err := pool.StartGame("foo", comm)
	assert.Equal(t, 1, len(pool.games))
	assert.Nil(t, err)

	comm2 := &TestCommunicator{}
	_, err = pool.JoinGame("foo", "bar", comm2)
	assert.Nil(t, err)

	assert.Equal(t, "bar", comm.NotifyArcherStateArcher.GetName())
	assert.Equal(t, "JOINED", comm.NotifyArcherStateState)
}

type TestCommunicator struct{
	NotifyArcherStateState string
	NotifyArcherStateArcher interfaces.Actor
}

func (c *TestCommunicator) NotifyMove(actor interfaces.Actor, x, y int) {
	panic("implement me")
}

func (c *TestCommunicator) NotifyShot(actor interfaces.Actor, target interfaces.Actor) {
	panic("implement me")
}

func (c *TestCommunicator) NotifyWon(actor interfaces.Actor) {
	panic("implement me")
}

func (c *TestCommunicator) NotifyFail(actor interfaces.Actor) {
	panic("implement me")
}

func (c *TestCommunicator) NotifyArcherState(state string, actor interfaces.Actor) {
	c.NotifyArcherStateState = state
	c.NotifyArcherStateArcher = actor
}
