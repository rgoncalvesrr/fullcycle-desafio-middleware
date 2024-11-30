package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/rgoncalvesrr/fullcycle-desafio-middleware/application"
	"github.com/rgoncalvesrr/fullcycle-desafio-middleware/configs"
	"log"
	"time"
)

type LimitRepositoryRedis struct {
	db *redis.Client
}

type LimiterTokenModel struct {
	Id            string    `json:"id"`
	LastRequestAt time.Time `json:"lastRequestAt"`
	BlockedUntil  time.Time `json:"blockedUntil"`
	Requests      int       `json:"requests"`
}

func NewRedisLimitRepository() *LimitRepositoryRedis {
	return &LimitRepositoryRedis{
		db: redis.NewClient(&redis.Options{
			Addr:     configs.Configs.CacheDbUrl,
			Password: configs.Configs.CacheDbPassword,
			DB:       0,
		}),
	}
}
func (r *LimitRepositoryRedis) Push(ctx context.Context, token *application.LimiterToken) error {
	content, err := json.Marshal(&LimiterTokenModel{
		Id:            token.ID,
		Requests:      token.GetRequests(),
		BlockedUntil:  token.GetBlockedUntil(),
		LastRequestAt: token.GetLastRequestAt(),
	})

	log.Printf("Objeto armazenado %s ", string(content))

	if err != nil {
		return err
	}
	if err := r.db.Set(ctx, token.ID, content, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r *LimitRepositoryRedis) Pull(ctx context.Context, keyToken string) (*application.LimiterToken, error) {
	content, e := r.db.Get(ctx, keyToken).Result()
	if e != nil {
		if e == redis.Nil {
			return nil, nil
		}
		fmt.Printf("erro a acessar redis: %v\n", e)
		return nil, e
	}

	var result *LimiterTokenModel

	e = json.Unmarshal([]byte(content), &result)
	if e != nil {
		return nil, e
	}

	log.Printf("Objeto recuperado %s ", string(content))

	return application.RestoreLimiterToken(
		result.Id,
		result.LastRequestAt,
		result.Requests,
		result.BlockedUntil,
	), nil
}
