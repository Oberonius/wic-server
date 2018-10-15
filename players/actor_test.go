package players

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"wic-server/interfaces"
)

func TestItSetsSpawnedStateOnSpawn(t *testing.T) {
	a := &Actor{}
	a.Spawn()
	assert.Equal(t, interfaces.ActorStateSpawned, a.GetState())
}

func TestItDoNothingOnShoot(t *testing.T) {
	a := &Actor{}
	a.Shoot(1, 1)
}

func TestItCanLeaveWithoutWorld(t *testing.T) {
	a := &Actor{}
	a.Leave()
}
