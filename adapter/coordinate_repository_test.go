package adapter_test

import (
	"context"
	"github.com/rgoncalvesrr/fullcycle-desafio-middleware/adapter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCoordinateRepository(t *testing.T) {
	r := adapter.NewCoordinateRepository()
	c, e := r.GetByCep(context.Background(), "09130220")

	assert.NotNil(t, r)
	assert.Nil(t, e)
	assert.NotNil(t, c)

}
