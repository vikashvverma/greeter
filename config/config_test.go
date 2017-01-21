package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig("config.json")
	require.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "20:02", config.Time)
	assert.Equal(t, "Happy B'day!", config.Subject)
}
func TestReadConfigFails(t *testing.T) {
	config, err := ReadConfig("/dev/void")
	assert.Nil(t, config)
	assert.Error(t, err, "could not read config file: %s")
}
