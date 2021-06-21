package _603_design_parking_system

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestParkingSystem(t *testing.T) {
	obj := Constructor(1, 1, 0)
	assert.Check(t, is.Equal(obj.AddCar(1), true))
	assert.Check(t, is.Equal(obj.AddCar(2), true))
	assert.Check(t, is.Equal(obj.AddCar(3), false))
	assert.Check(t, is.Equal(obj.AddCar(1), false))
}
