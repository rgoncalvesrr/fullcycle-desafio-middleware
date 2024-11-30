package application

import (
	"errors"
	"time"
)

var (
	ErrRateLimitExceeded = errors.New("rate limit exceeded")
	ErrBlockedKeyOrIP    = errors.New("blocked key or ip address")
)

type LimiterToken struct {
	ID                string
	lastRequestsReset time.Time
	requests          int
	blockedUntil      time.Time
}

func NewLimiterToken(id string) *LimiterToken {
	return &LimiterToken{
		ID:                id,
		requests:          1,
		blockedUntil:      time.Now().Add(-1 * time.Second),
		lastRequestsReset: time.Now(),
	}
}

func RestoreLimiterToken(
	id string,
	lastRequestsReset time.Time,
	requests int,
	blockedUntil time.Time) *LimiterToken {

	return &LimiterToken{
		ID:                id,
		requests:          requests,
		blockedUntil:      blockedUntil,
		lastRequestsReset: lastRequestsReset,
	}
}

func (l *LimiterToken) Request(limitPerSecond int, blockedDuration time.Duration) error {
	// Impede requisição para o token bloqueado
	if l.IsBlocked() {
		return ErrBlockedKeyOrIP
	}

	// Reinicia contador de requisições
	if l.lastRequestsReset.Add(time.Second).Before(time.Now()) {
		l.requests = 1
		l.lastRequestsReset = time.Now()
		return nil
	}

	// Incrementa contador de requisições e verifica se houve extrapolação
	l.requests++
	if l.requests > limitPerSecond {
		l.blockedUntil = time.Now().Add(blockedDuration)
		return ErrRateLimitExceeded
	}

	return nil
}

func (l *LimiterToken) IsBlocked() bool {
	return l.blockedUntil.After(time.Now())
}
func (l *LimiterToken) GetLastRequestAt() time.Time {
	return l.lastRequestsReset
}
func (l *LimiterToken) GetBlockedUntil() time.Time {
	return l.blockedUntil
}
func (l *LimiterToken) GetRequests() int {
	return l.requests
}
