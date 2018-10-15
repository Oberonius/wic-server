package interfaces

type Communicator interface {
	NotifyMove(actor Actor, x, y int)
	NotifyShot(actor Actor, target Actor)
	NotifyWon(actor Actor)
	NotifyFail(actor Actor)
	NotifyArcherState(state string, actor Actor)
}
