package application

import "context"

type ICoordinateRepository interface {
	GetByCep(ctx context.Context, cep string) (*Coordinate, error)
}

type IWeatherRepository interface {
	GetTemperature(ctx context.Context, coordinate *Coordinate) (*Weather, error)
}

type ILimiterCacheRepository interface {
	Push(ctx context.Context, token *LimiterToken) error
	Pull(ctx context.Context, keyToken string) (*LimiterToken, error)
}
