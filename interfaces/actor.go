package interfaces

const (
	_                 = iota
	ActorStateSpawned
	ActorStateLeaved
)

const (
	ActorRoleZombie = iota
	ActorRoleArcher
)

type Actor interface {
	Spawn()
	Stop()
	GetName() string
	IsInPosition(x, y int) bool
	GetRole() int
	GetState() int

	Shoot(int, int)
	Leave()
}
