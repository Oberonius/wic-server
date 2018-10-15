package interfaces

import "github.com/stretchr/testify/mock"

type MockCommunicator struct {
	mock.Mock
}

func (c *MockCommunicator) NotifyMove(actor Actor, x, y int) {
	c.Called(actor, x, y)
}

func (c *MockCommunicator) NotifyShot(actor Actor, target Actor) {
	c.Called(actor, target)
}

func (c *MockCommunicator) NotifyWon(actor Actor) {
	c.Called(actor)
}

func (c *MockCommunicator) NotifyFail(actor Actor) {
	c.Called(actor)
}

func (c *MockCommunicator) NotifyArcherState(state string, actor Actor) {
	c.Called(state, actor)
}
