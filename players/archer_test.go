package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"wic-server/interfaces"
)

func TestItCreatesArcher(t *testing.T) {
	a := NewArcher("foo", nil)
	assert.Equal(t, "foo", a.GetName())
}

func TestItNotifiesWorldAboutShooting(t *testing.T) {
	w := &interfaces.MockWorld{}
	a := NewArcher("foo", w)
	w.On("Shoot", a, 1, 2).Once()
	a.Shoot(1, 2)
}

func TestItWillReturnOnShootingWithoutWorld(t *testing.T){
	a := NewArcher("foo", nil)
	a.Shoot(1, 2)
}

func TestItReturnProperRole(t *testing.T){
	a := NewArcher("foo", nil)
	assert.Equal(t, interfaces.ActorRoleArcher, a.GetRole())
}