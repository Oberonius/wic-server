package game

import "wic-server/interfaces"

func (b *Board) Shoot(shooter interfaces.Actor, x, y int) {
	b.Lock()

	shouldStop := false
	zombies, killed := b.processShooting(shooter, x, y)

	b.notifyAll(func(c interfaces.Communicator) {
		c.NotifyShot(shooter, killed)
	})

	if zombies == 0 {
		b.notifyAll(func(c interfaces.Communicator) {
			c.NotifyWon(shooter)
		})
		shouldStop = true
	}
	b.Unlock()

	if shouldStop {
		b.Stop()
	}
}

func (b *Board) Move(person interfaces.Actor, x, y int) {
	b.Lock()

	shouldStop := false

	b.notifyAll(func(c interfaces.Communicator) {
		c.NotifyMove(person, x, y)
	})

	if person.GetRole() == interfaces.ActorRoleZombie && y == maxY {
		b.notifyAll(func(c interfaces.Communicator) {
			c.NotifyFail(person)
		})
		shouldStop = true
	}

	b.Unlock()

	if shouldStop {
		b.Stop()
	}
}

func (b *Board) processShooting(shooter interfaces.Actor, x int, y int) (int, interfaces.Actor) {
	zombies := 0
	var killed interfaces.Actor
	for n, a := range b.actors {
		if n != shooter.GetName() && a.IsInPosition(x, y) {
			a.Stop()
			delete(b.actors, n)
			killed = a
			continue
		}

		if a.GetRole() == interfaces.ActorRoleZombie {
			zombies++
		}
	}

	return zombies, killed
}
