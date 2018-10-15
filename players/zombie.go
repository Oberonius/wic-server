package players

import (
	"math/rand"
	"time"
	"wic-server/interfaces"
)

const (
	moveByX = 0
	moveByY = 1
)

//Zombie struct represents zombie actor type
type Zombie struct {
	Actor
	moveDelay    time.Duration
	moveMeCh     chan bool
	axesRnd      func(int) int
	directionRnd func(int) int
}

//NewZombie creates new named zombie in desired coordinates of the world
func NewZombie(x, y int, name string, world interfaces.World) *Zombie {
	return &Zombie{
		Actor: Actor{
			role:    interfaces.ActorRoleZombie,
			name:    name,
			x:       x,
			y:       y,
			world:   world,
			leaveCh: make(chan bool, 1),
		},
		moveDelay:    2000 * time.Millisecond,
		moveMeCh:     make(chan bool, 1),
		axesRnd:      rand.Intn,
		directionRnd: rand.Intn,
	}
}

//Spawn starts all the lifecycle
func (z *Zombie) Spawn() {
	go z.doWalking()
	go z.listen()
	z.Actor.Spawn()
}

//Leave stops zombie and do some cleaning
func (z *Zombie) Leave() {
	z.Actor.Leave()
	close(z.moveMeCh)
}

//doWalking makes smooth moving possible.
//it creates move event every two seconds
func (z *Zombie) doWalking() {
	for {
		time.Sleep(z.moveDelay)
		z.RLock()
		if z.world == nil || z.state == interfaces.ActorStateLeaved {
			z.RUnlock()
			return
		}
		z.RUnlock()
		z.moveMeCh <- true
	}
}

//listen is the main event loop
func (z *Zombie) listen() {
	for {
		select {
		case <-z.leaveCh:
			return
		case <-z.moveMeCh:
			z.move()
		}
	}
}

//move changes zombie coordinates and notifies the world on this event
func (z *Zombie) move() {
	z.Lock()
	if world := z.world; world != nil {
		z.moveRandomly()
		x, y := z.x, z.y
		z.Unlock()
		world.Move(z, x, y)
	} else {
		z.Unlock()
	}
}

//moveRandomly changes zombie coordinates with one step by X or Y axe
//zombie goes forward and can do steps left or right
func (z *Zombie) moveRandomly() {
	if z.world == nil {
		return
	}

	if z.axesRnd(2) == moveByX {
		direction := z.directionRnd(2)
		if direction == 0 {
			direction = -1
		}
		z.moveByX(direction)
	} else {
		z.moveByY()
	}
}

func (z *Zombie) moveByX(direction int) {
	if z.x+direction > z.world.GetMaxX() || z.x+direction < 0 {
		z.x -= direction
	} else {
		z.x += direction
	}
}

func (z *Zombie) moveByY() {
	if z.y+1 > z.world.GetMaxY() {
		z.y--
	} else {
		z.y++
	}
}
