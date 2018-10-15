package interfaces

type World interface {
	Shoot(Actor, int, int)
	Move(Actor, int, int)
	NotifyLeaved(Actor)
	GetMaxX() int
	GetMaxY() int
}
