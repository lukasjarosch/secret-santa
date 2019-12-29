package middleware

import (
	"net/http"
	"time"

	"github.com/lukasjarosch/genki/logger"
)

func Logging(handler http.Handler, skipEndpoint ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := statusWriter{ResponseWriter: w}

		for _, skip := range skipEndpoint {
			if r.URL.String() == skip {
				handler.ServeHTTP(&sw, r)
				return
			}
		}

		log := logger.WithMetadata(r.Context())
		log = log.WithFields(logger.Fields{
			"req.method": r.Method,
			"req.url":    r.URL,
		})
		log.Infof("incoming request to %s %s", r.Method, r.URL)
		defer func(started time.Time) {
			log = log.WithFields(logger.Fields{
				"took":   time.Since(started),
				"status": sw.status,
			})
			log.Infof("served request to %s", r.URL)
		}(time.Now())
		handler.ServeHTTP(&sw, r)
	})
}

func LoggingFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		sw := statusWriter{ResponseWriter: writer}

		log := logger.WithMetadata(request.Context())
		log = log.WithFields(logger.Fields{
			"req.method": request.Method,
			"req.url":    request.URL,
		})
		log.Infof("incoming request to %s %s", request.Method, request.URL)
		defer func(started time.Time) {
			log = log.WithFields(logger.Fields{
				"took":   time.Since(started),
				"status": sw.status,
			})
			log.Infof("served request to %s", request.URL)
		}(time.Now())
		next.ServeHTTP(&sw, request)
	}
}
