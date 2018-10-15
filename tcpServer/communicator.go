package tcpServer

import (
	"fmt"
	"io"
	"wic-server/interfaces"
)

type StdCommunicator struct {
	writer io.Writer
}

func NewCommunicator(writer io.Writer) *StdCommunicator {
	return &StdCommunicator{
		writer: writer,
	}
}

func (sc *StdCommunicator) NotifyMove(actor interfaces.Actor, x, y int) {
	fmt.Fprintf(sc.writer, "WALK %s %v %v\n", actor.GetName(), x, y)
}

func (sc *StdCommunicator) NotifyArcherState(state string, actor interfaces.Actor) {
	fmt.Fprintf(sc.writer, "%s %s\n", state, actor.GetName())
}

func (sc *StdCommunicator) NotifyShot(actor interfaces.Actor, target interfaces.Actor) {
	if target == nil {
		fmt.Fprintf(sc.writer, "BOOM %s 0\n\n", actor.GetName())
	} else {
		fmt.Fprintf(sc.writer, "BOOM %s 1 %s\n\n", actor.GetName(), target.GetName())
	}
}

func (sc *StdCommunicator) NotifyWon(actor interfaces.Actor) {
	fmt.Fprintf(sc.writer, "WINNER %s\n\n", actor.GetName())
}

func (sc *StdCommunicator) NotifyFail(actor interfaces.Actor) {
	fmt.Fprintf(sc.writer, "ZOMBIE-WINNER %s\n\n", actor.GetName())
}
