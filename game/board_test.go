package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"wic-server/interfaces"
)

func TestItCreatesBoard(t *testing.T) {
	b := NewBoard("foo", func(*Board) {})
	assert.Equal(t, "foo", b.GetGameName())
	assert.False(t, b.IsInGame())
}

func createBoardWithTwoPlayers() (
	*Board,
	*interfaces.MockActor, *interfaces.MockCommunicator,
	*interfaces.MockActor, *interfaces.MockCommunicator,
) {
	b := NewBoard("foo", func(b *Board) {})

	arBody := &interfaces.MockActor{}
	arBody.On("GetName").Return("foo")
	arComm := &interfaces.MockCommunicator{}
	b.AddPlayer(arBody, arComm)

	ar2Body := &interfaces.MockActor{}
	ar2Body.On("GetName").Return("foo2")
	ar2Comm := &interfaces.MockCommunicator{}
	b.AddPlayer(ar2Body, ar2Comm)

	return b, arBody, arComm, ar2Body, ar2Comm
}

func TestItStartsGame(t *testing.T) {
	b := NewBoard("foo", func(b *Board) { assert.True(t, b.IsInGame()) })

	arBody := &interfaces.MockActor{}
	arBody.On("GetName").Return("foo")
	arBody.On("Spawn").Once()
	arComm := &interfaces.MockCommunicator{}

	b.AddPlayer(arBody, arComm)
	b.Run()
	assert.True(t, b.IsInGame())
	b.Run()
}

func TestItNotifiesOnNewPlayer(t *testing.T) {
	b := NewBoard("foo", func(b *Board) { assert.True(t, b.IsInGame()) })

	arBody := &interfaces.MockActor{}
	arBody.On("GetName").Return("foo")
	arBody.On("Spawn").Once()
	arComm := &interfaces.MockCommunicator{}

	b.AddPlayer(arBody, arComm)
	b.Run()

	ar2Body := &interfaces.MockActor{}
	ar2Body.On("GetName").Return("foo2")
	ar2Body.On("Spawn").Once()
	ar2Comm := &interfaces.MockCommunicator{}

	arComm.On("NotifyArcherState", archerJoined, ar2Body).Once()
	b.AddPlayer(ar2Body, ar2Comm)
}

func TestItDoesNotAddDuplicatedPlayerNames(t *testing.T) {
	b := NewBoard("foo", func(b *Board) { assert.True(t, b.IsInGame()) })

	arBody := &interfaces.MockActor{}
	arBody.On("GetName").Return("foo")
	arBody.On("Spawn").Once()
	arComm := &interfaces.MockCommunicator{}

	b.AddPlayer(arBody, arComm)
	b.Run()

	ar2Body := &interfaces.MockActor{}
	ar2Body.On("GetName").Return("foo")
	ar2Body.On("Spawn").Once()
	ar2Comm := &interfaces.MockCommunicator{}

	assert.NotNil(t, b.AddPlayer(ar2Body, ar2Comm))
}

func TestItStopsTheGame(t *testing.T) {
	b, a1, _, a2, _ := createBoardWithTwoPlayers()

	a1.On("Spawn").Once()
	a2.On("Spawn").Once()
	b.Run()
	assert.True(t, b.IsInGame())

	b.stateChangeObserver = func(b *Board) { assert.False(t, b.IsInGame()) }
	a1.On("Stop").Once()
	a2.On("Stop").Once()
	b.Stop()
	assert.False(t, b.IsInGame())
}

func TestItProperlyWorkWithLeaving(t *testing.T) {
	b, a1, _, a2, c2 := createBoardWithTwoPlayers()

	a1.On("Spawn").Once()
	a1.On("GetRole").Return(interfaces.ActorRoleArcher)
	a2.On("Spawn").Once()
	a2.On("GetRole").Return(interfaces.ActorRoleArcher)
	b.Run()
	assert.True(t, b.IsInGame())

	c2.On("NotifyArcherState", archerLeaved, a1).Once()
	b.NotifyLeaved(a1)
	b.NotifyLeaved(a2)
	assert.False(t, b.IsInGame())
}

func TestItDoesNotNotifyOnZombieLeaving(t *testing.T) {
	b, a1, _, a2, _ := createBoardWithTwoPlayers()

	a1.On("Spawn").Once()
	a1.On("GetRole").Return(interfaces.ActorRoleZombie)
	a2.On("Spawn").Once()
	a2.On("GetRole").Return(interfaces.ActorRoleArcher)
	b.Run()
	assert.True(t, b.IsInGame())

	b.NotifyLeaved(a1)
}

