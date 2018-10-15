package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"wic-server/interfaces"
	"time"
)

var (
	xRnd   = func(a int) int { return moveByX }
	yRnd   = func(a int) int { return moveByY }
	posRnd = func(a int) int { return 1 }
	negRnd = func(a int) int { return 0 }
)

func newZombieWithWorld(x, y int) *Zombie {
	w := &interfaces.MockWorld{}
	w.On("GetMaxX").Return(10)
	w.On("GetMaxY").Return(10)
	return NewZombie(x, y, "DK", w)
}

func TestZombieLifecycle(t *testing.T) {
	w := &interfaces.MockWorld{}
	z := NewZombie(2, 2, "DK", w)

	w.On("GetMaxX").Return(10)
	w.On("GetMaxY").Return(10)
	w.On("Move", z, 3, 2).Once()
	w.On("NotifyLeaved", &z.Actor).Once()

	z.axesRnd = xRnd
	z.directionRnd = posRnd
	z.moveDelay = 50 * time.Millisecond
	z.Spawn()

	//move zombie in one position
	time.Sleep(60 * time.Millisecond)
	assert.True(t, z.IsInPosition(3, 2))

	//test leave functionality
	z.Leave()
	time.Sleep(60 * time.Millisecond)
}

func TestItWillNotMoveWithoutWorld(t *testing.T) {
	z := NewZombie(3, 3, "DK", nil)
	z.axesRnd = xRnd
	z.directionRnd = posRnd
	z.moveRandomly()
	assert.Equal(t, 3, z.x)
	assert.Equal(t, 3, z.y)

}

func TestItAnswerFalseOnPositionsWithoutWorld(t *testing.T) {
	z := NewZombie(3, 3, "DK", nil)
	assert.False(t, z.IsInPosition(3, 3))
}

func TestItStoresPosition(t *testing.T) {
	z := newZombieWithWorld(3, 4)
	assert.True(t, z.IsInPosition(3, 4))
	assert.False(t, z.IsInPosition(1, 5))
}

func TestItMovesByX(t *testing.T) {
	z := newZombieWithWorld(3, 4)
	z.axesRnd = xRnd

	z.directionRnd = posRnd
	z.moveRandomly()
	assert.True(t, z.IsInPosition(4, 4))

	z.directionRnd = negRnd
	z.moveRandomly()
	assert.True(t, z.IsInPosition(3, 4))
}

func TestItBumpsOnTheWallsByX(t *testing.T) {
	z := newZombieWithWorld(0, 4)
	z.axesRnd = xRnd
	z.directionRnd = negRnd
	z.moveRandomly()
	assert.True(t, z.IsInPosition(1, 4))

	z = newZombieWithWorld(10, 4)
	z.axesRnd = xRnd
	z.directionRnd = posRnd
	z.moveRandomly()
	assert.True(t, z.IsInPosition(9, 4))
}

func TestItAlwaysMovesForwardByY(t *testing.T) {
	z := newZombieWithWorld(3, 4)
	z.axesRnd = yRnd
	z.directionRnd = posRnd
	z.moveRandomly()
	assert.True(t, z.IsInPosition(3, 5))

	z.directionRnd = negRnd
	z.moveRandomly()
	assert.True(t, z.IsInPosition(3, 6))
}

func TestItBumpsOnTheWallsByY(t *testing.T) {
	z := newZombieWithWorld(5, 30)
	z.axesRnd = yRnd

	z.moveRandomly()
	assert.True(t, z.IsInPosition(5, 29))
}

func TestItReturnZombieRole(t *testing.T) {
	a := NewZombie(1, 1, "foo", nil)
	assert.Equal(t, interfaces.ActorRoleZombie, a.GetRole())
}
