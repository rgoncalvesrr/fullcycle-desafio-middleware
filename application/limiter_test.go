package application_test

import (
	"github.com/rgoncalvesrr/fullcycle-desafio-middleware/application"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewLimiterToken(t *testing.T) {
	expectedId := "123"
	token := application.NewLimiterToken(expectedId)

	assert.NotNil(t, token)
	assert.Equal(t, expectedId, token.ID)
}

func TestLimiterToken_Valid_Request(t *testing.T) {
	expectedId := "123"
	token := application.NewLimiterToken(expectedId)
	e := token.Request(500, time.Second)
	assert.Nil(t, e)
}

func TestLimiterToken_RateLimitExceeded(t *testing.T) {
	expectedId := "123"
	limitPerSec := 1
	token := application.NewLimiterToken(expectedId) // A criação do token conta como primeira requisição
	e := token.Request(limitPerSec, time.Second)     // Segunda requisição
	assert.ErrorIs(t, e, application.ErrRateLimitExceeded)
}

func TestLimiterToken_BlockedKeyOrIP(t *testing.T) {
	expectedId := "123"
	limitPerSec := 1
	token := application.NewLimiterToken(expectedId) // A criação do token conta como primeira requisição
	_ = token.Request(limitPerSec, time.Second)      // Segunda requisição -> bloqueará por excesso de requisições
	e := token.Request(limitPerSec, time.Second)     // Terceira requisição
	assert.ErrorIs(t, e, application.ErrBlockedKeyOrIP)
}

func TestLimiterToken_ElapsedTimeBlockedKeyOrIP(t *testing.T) {
	expectedId := "123"
	limitPerSec := 1
	blockedDuration := time.Second
	token := application.NewLimiterToken(expectedId) // A criação do token conta como primeira requisição

	e := token.Request(limitPerSec, blockedDuration) // Segunda requisição -> bloqueada por excesso de requisições
	assert.ErrorIs(t, e, application.ErrRateLimitExceeded)

	e = token.Request(limitPerSec, blockedDuration) // Terceira requisição -> bloquada por tempo
	assert.ErrorIs(t, e, application.ErrBlockedKeyOrIP)
	assert.True(t, token.IsBlocked())

	time.Sleep(blockedDuration)
	e = token.Request(limitPerSec, blockedDuration) // Quarta requisição -> penalidade cumprida
	assert.Nil(t, e)
	assert.False(t, token.IsBlocked())
}
