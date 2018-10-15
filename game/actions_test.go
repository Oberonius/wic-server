package game

import (
	"testing"
	"wic-server/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestItKillsZombie(t *testing.T) {
	b, a1, _, a2, c2 := createBoardWithTwoPlayers()
	a1.On("Spawn").Once()
	a1.On("GetRole").Return(interfaces.ActorRoleZombie)
	a1.On("IsInPosition", 5, 5).Return(true)

	a2.On("Spawn").Once()
	a2.On("GetRole").Return(interfaces.ActorRoleArcher)
	b.Run()

	a1.On("Stop").Once()
	a2.On("Stop").Once()
	c2.On("NotifyShot", a2, a1).Once()
	c2.On("NotifyWon", a2).Once()
	b.Shoot(a2, 5, 5)

	assert.False(t, b.IsInGame())
}

func TestItJustNotifiesIfArcherMissed(t *testing.T) {
	b, a1, c1, a2, c2 := createBoardWithTwoPlayers()
	a1.On("Spawn").Once()
	a1.On("GetRole").Return(interfaces.ActorRoleZombie)
	a1.On("IsInPosition", 5, 5).Return(false)

	a2.On("Spawn").Once()
	a2.On("GetRole").Return(interfaces.ActorRoleArcher)
	b.Run()

	c2.On("NotifyShot", a2, nil).Once()
	c1.On("NotifyShot", a2, nil).Once()
	b.Shoot(a2, 5, 5)

	assert.True(t, b.IsInGame())
}

func TestItWillNotifyOnZombieMove(t *testing.T) {
	b, a1, c1, a2, c2 := createBoardWithTwoPlayers()
	a1.On("Spawn").Once()
	a1.On("GetRole").Return(interfaces.ActorRoleZombie)

	a2.On("Spawn").Once()
	a2.On("GetRole").Return(interfaces.ActorRoleArcher)
	b.Run()

	c2.On("NotifyMove", a1, 5, 5).Once()
	c1.On("NotifyMove", a1, 5, 5).Once()
	b.Move(a1, 5, 5)

	assert.True(t, b.IsInGame())
}

func TestItLetsZombieWin(t *testing.T) {
	b, a1, c1, a2, c2 := createBoardWithTwoPlayers()
	a1.On("Spawn").Once()
	a1.On("GetRole").Return(interfaces.ActorRoleZombie)

	a2.On("Spawn").Once()
	a2.On("GetRole").Return(interfaces.ActorRoleArcher)
	b.Run()

	//it notifies all on zombie move
	c1.On("NotifyMove", a1, 5, b.GetMaxY()).Once()
	c2.On("NotifyMove", a1, 5, b.GetMaxY()).Once()

	//it notifies all on zombie win
	c1.On("NotifyFail", a1).Once()
	c2.On("NotifyFail", a1).Once()

	//it stops the game
	a1.On("Stop").Once()
	a2.On("Stop").Once()

	b.Move(a1, 5, b.GetMaxY())

	assert.False(t, b.IsInGame())
}
