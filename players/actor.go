package players

import (
	"sync"
	"wic-server/interfaces"
)

//Actor is abstract acting person
type Actor struct {
	sync.RWMutex
	name    string
	world   interfaces.World
	x, y    int
	leaveCh chan bool
	role    int
	state   int
}

//Spawn function starts new actor lifecycle
func (a *Actor) Spawn() {
	a.Lock()
	defer a.Unlock()
	a.state = interfaces.ActorStateSpawned
}

//GetName returns actor's name
func (a *Actor) GetName() string {
	a.RLock()
	defer a.RUnlock()
	return a.name
}

//Shoot is a stub for future implementation in specific actor types
func (a *Actor) Shoot(x, y int) {
	//boom!
}

//IsInPosition matches provided coordinates with current actor position
func (a *Actor) IsInPosition(x, y int) bool {
	a.RLock()
	defer a.RUnlock()
	return a.world != nil && a.x == x && a.y == y
}

//Stop stops all actor activity
func (a *Actor) Stop() {
	a.Lock()
	defer a.Unlock()

	a.state = interfaces.ActorStateLeaved
	if a.leaveCh != nil {
		a.leaveCh <- true
		close(a.leaveCh)
	}
	if a.world != nil {
		a.world = nil
	}
}

//Leave notifies the world on actor leave event and stops all actor activity
func (a *Actor) Leave() {
	a.RLock()
	if world := a.world; world != nil {
		a.RUnlock()
		world.NotifyLeaved(a)
	} else {
		a.RUnlock()
	}

	a.Stop()
}

//GetRole returns current actor role (zombie/archer)
func (a *Actor) GetRole() int {
	a.RLock()
	defer a.RUnlock()
	return a.role
}

//GetState returns current actor state (leaved, spawned, ...)
func (a *Actor) GetState() int {
	a.RLock()
	defer a.RUnlock()
	return a.state
}
