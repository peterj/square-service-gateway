package server

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// UserAgentHeaderName is the User-Agent header name
	UserAgentHeaderName = "User-Agent"
)

// WithLogging middleware logs all requests
func WithLogging(next http.Handler) http.Handler {
	start := time.Now()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof(
			"%s | %s | %s | %s",
			r.Method,
			r.RequestURI,
			r.Header.Get(UserAgentHeaderName),
			time.Since(start),
		)
		next.ServeHTTP(w, r)
	})
}
