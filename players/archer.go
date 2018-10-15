package players

import "wic-server/interfaces"

//Archer struct represents archer actor type
type Archer struct {
	Actor
}

//NewArcher creates new named archer in the world
func NewArcher(name string, world interfaces.World) *Archer {
	return &Archer{
		Actor{
			role:  interfaces.ActorRoleArcher,
			name:  name,
			x:     -1,
			y:     -1,
			world: world,
		},
	}
}

//Shoot performs shooting. It just notifies the world about this event
func (a *Archer) Shoot(x, y int) {
	a.RLock()
	if world := a.world; world == nil {
		a.RUnlock()
		return
	} else {
		a.RUnlock()
		a.world.Shoot(a, x, y)
	}
}

