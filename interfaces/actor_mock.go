package interfaces

import "github.com/stretchr/testify/mock"

type MockActor struct {
	mock.Mock
}

func (a *MockActor) Spawn() {
	a.Mock.Called()
}

func (a *MockActor) Stop() {
	a.Mock.Called()
}

func (a *MockActor) GetName() string {
	args := a.Mock.Called()
	return args.String(0)
}

func (a *MockActor) IsInPosition(x, y int) bool {
	args := a.Mock.Called(x, y)
	return args.Bool(0)
}

func (a *MockActor) GetRole() int {
	args := a.Mock.Called()
	return args.Int(0)
}

func (a *MockActor) GetState() int {
	args := a.Mock.Called()
	return args.Int(0)
}

func (a *MockActor) Shoot(x, y int) {
	a.Mock.Called(x, y)
}

func (a *MockActor) Leave() {
	a.Mock.Called()
}
