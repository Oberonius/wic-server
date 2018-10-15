package interfaces

import (
	"github.com/stretchr/testify/mock"
)

type MockWorld struct {
	mock.Mock
}

func (w *MockWorld) Shoot(a Actor, x int, y int) {
	w.Mock.Called(a, x, y)
}

func (w *MockWorld) Move(a Actor, x int, y int) {
	w.Mock.Called(a, x, y)
}

func (w *MockWorld) NotifyLeaved(a Actor) {
	w.Mock.Called(a)
}

func (w *MockWorld) GetMaxX() int {
	args := w.Mock.Called()
	return args.Int(0)
}

func (w *MockWorld) GetMaxY() int {
	args := w.Mock.Called()
	return args.Int(0)
}
