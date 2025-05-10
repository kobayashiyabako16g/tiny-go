package middleware

import (
	"net/http"
	"time"

	"github.com/kobayashiyabako16g/tiny-go/pkg/logger"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(lrw, r)

		logger.Info(ctx, "HTTP Request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", lrw.statusCode,
			"duration", time.Since(start),
			"user_agent", r.UserAgent(),
			"ip", r.RemoteAddr,
		)
	})
}

// loggingResponseWriter はレスポンスのステータスコードを取得するためのラッパー
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
