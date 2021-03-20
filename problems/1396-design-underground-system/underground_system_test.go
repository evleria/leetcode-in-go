package _396_design_underground_system

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestUndergroundSystem(t *testing.T) {
	us := Constructor()
	us.CheckIn(45, "Leyton", 3)
	us.CheckIn(32, "Paradise", 8)
	us.CheckIn(27, "Leyton", 10)
	us.CheckOut(45, "Waterloo", 15)
	us.CheckOut(27, "Waterloo", 20)
	us.CheckOut(32, "Cambridge", 22)
	assert.Check(t, is.Equal(us.GetAverageTime("Paradise", "Cambridge"), 14.0))
	assert.Check(t, is.Equal(us.GetAverageTime("Leyton", "Waterloo"), 11.0))
	us.CheckIn(10, "Leyton", 24)
	assert.Check(t, is.Equal(us.GetAverageTime("Leyton", "Waterloo"), 11.0))
	us.CheckOut(10, "Waterloo", 38)
	assert.Check(t, is.Equal(us.GetAverageTime("Leyton", "Waterloo"), 12.0))
}
