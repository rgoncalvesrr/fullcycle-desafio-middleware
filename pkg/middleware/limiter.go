package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/rgoncalvesrr/fullcycle-desafio-middleware/adapter"
	"github.com/rgoncalvesrr/fullcycle-desafio-middleware/application"
	"github.com/rgoncalvesrr/fullcycle-desafio-middleware/configs"
	"net"
	"net/http"
)

type outBlock struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Limiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ti := application.LimiterInputToken{
			Key:                  r.Header.Get("API_KEY"),
			IP:                   ip,
			RequestsPerSecondKey: configs.Configs.LimiterReqPerSecKey,
			RequestsPerSecondIp:  configs.Configs.LimiterReqPerSecIP,
			RequestPenalty:       configs.Configs.LimiterPenaltySec,
		}

		limiterService := application.NewLimiterService(adapter.NewRedisLimitRepository())

		w.Header().Add("Content-Type", "application/json")

		// Verifica bloqueio
		if limiterService.RegisterRequest(ctx, ti) {
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {

			w.WriteHeader(http.StatusTooManyRequests)

			err := json.NewEncoder(w).Encode(&outBlock{
				Status:  http.StatusTooManyRequests,
				Message: "you have reached the maximum number of requests or actions allowed within a certain time frame",
			})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})
}
