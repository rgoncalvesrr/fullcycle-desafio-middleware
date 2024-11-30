package configs_test

import (
	"github.com/rgoncalvesrr/fullcycle-desafio-middleware/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := configs.Init("..")
	assert.NotNil(t, c)
}
