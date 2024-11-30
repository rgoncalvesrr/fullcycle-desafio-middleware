package application

import (
	"context"
	"fmt"
	"time"
)

type LimiterService struct {
	repository ILimiterCacheRepository
}

type LimiterInputToken struct {
	IP                   string
	Key                  string
	RequestsPerSecondKey int
	RequestsPerSecondIp  int
	RequestPenalty       int
}

func NewLimiterService(repository ILimiterCacheRepository) *LimiterService {
	return &LimiterService{repository: repository}
}

func (l *LimiterService) IsBlocked(ctx context.Context, token string) bool {
	t, e := l.repository.Pull(ctx, token)

	if e != nil {
		return false
	}

	if t != nil {
		return t.IsBlocked()
	}
	return false
}

func (l *LimiterService) RegisterRequest(ctx context.Context, token LimiterInputToken) bool {
	var tokenKey string
	var requestsPerSecond int

	if token.Key != "" {
		if l.IsBlocked(ctx, token.IP) {
			return false
		}

		tokenKey = token.Key
		requestsPerSecond = token.RequestsPerSecondKey
	} else {
		tokenKey = token.IP
		requestsPerSecond = token.RequestsPerSecondIp
	}

	t, e := l.repository.Pull(ctx, tokenKey)

	if e != nil {
		fmt.Printf("request for token %s failed: %w", tokenKey, e)
		return false
	}

	if t == nil {
		t = NewLimiterToken(tokenKey)
	} else {
		blockedUntil := time.Second * time.Duration(token.RequestPenalty)
		e = t.Request(requestsPerSecond, blockedUntil)
	}
	if e != nil {
		fmt.Printf("request for token %s failed: %w", tokenKey, e)
		return false
	}
	e = l.repository.Push(ctx, t)

	return e == nil
}
