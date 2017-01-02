package job

import (
	"testing"
	"github.com/vikashvverma/greeter/person"
	"github.com/stretchr/testify/assert"
)

func TestGreetingFails(t *testing.T) {
	p := person.Person{Name:"Vikash"}
	c, err := greeting(p)
	assert.Nil(t, err)
	assert.Empty(t, c.Value)
}
