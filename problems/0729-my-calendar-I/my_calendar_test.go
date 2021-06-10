package _729_my_calendar_I

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMyCalendar(t *testing.T) {
	calendar := Constructor()
	assert.Check(t, is.Equal(calendar.Book(10, 20), true))
	assert.Check(t, is.Equal(calendar.Book(15, 25), false))
	assert.Check(t, is.Equal(calendar.Book(20, 30), true))

	calendar = Constructor()
	assert.Check(t, is.Equal(calendar.Book(47, 50), true))
	assert.Check(t, is.Equal(calendar.Book(33, 41), true))
	assert.Check(t, is.Equal(calendar.Book(39, 45), false))
	assert.Check(t, is.Equal(calendar.Book(33, 42), false))
	assert.Check(t, is.Equal(calendar.Book(25, 32), true))
	assert.Check(t, is.Equal(calendar.Book(26, 35), false))
	assert.Check(t, is.Equal(calendar.Book(19, 25), true))
	assert.Check(t, is.Equal(calendar.Book(3, 8), true))
	assert.Check(t, is.Equal(calendar.Book(8, 13), true))
	assert.Check(t, is.Equal(calendar.Book(18, 27), false))
}
