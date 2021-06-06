package middleware

import (
	"net/http"
	"time"

	"github.com/gopher-translator-service/pkg/logger"

	chi_middleware "github.com/go-chi/chi/middleware"
)

func LogIncommingReq(log logger.Logger) func(next http.Handler) http.Handler {
	ret := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ctx := r.Context()
			requestId := ctx.Value(chi_middleware.RequestIDKey)
			if requestId == nil {
				requestId = "0000-0000-0000-0000"
			}

			log.Infof("RequstId: %s Started: %v Method: %s URL: %s\n",
				requestId, start, r.Method, r.URL.String())

			defer func () {
				log.Infof("RequstId: %s Took: %v Method: %s URL: %s\n",
					requestId, time.Since(start), r.Method, r.URL.String())
			}()

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}

	return ret
}
