package game

import "wic-server/interfaces"

func (b *Board) notifyAll(cb func(interfaces.Communicator)) {
	b.notifyExcluding(cb, "")
}

func (b *Board) notifyExcluding(cb func(interfaces.Communicator), excludeName string) {
	for n := range b.actors {
		if cc := b.commChannels[n]; n != excludeName && cc != nil {
			cb(cc)
		}
	}
}
