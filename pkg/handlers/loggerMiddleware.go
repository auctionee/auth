package handlers

import (
	"context"
	"github.com/auctionee/auth/pkg/logger"
	"net/http"
)

func LoggerMiddleware(ctx context.Context, next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logger.GetLogger(ctx)
		l.Info("new request", r.URL)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}