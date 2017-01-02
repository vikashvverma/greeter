package person

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsToday(t *testing.T) {
	assert.False(t, (&Person{Month: time.January, Day: 1}).IsToday())
	assert.True(t, (&Person{Month: time.Now().Month(), Day: time.Now().Day()}).IsToday())
	assert.False(t, (&Person{Month: time.February, Day: 2}).IsToday())
}
