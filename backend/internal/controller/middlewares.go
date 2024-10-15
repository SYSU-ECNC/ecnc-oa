package controller

import (
	"net/http"
	"time"
)

func (ctrl *Controller) log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		clientIP := r.RemoteAddr
		method := r.Method
		url := r.URL.String()

		next.ServeHTTP(w, r)

		duration := time.Since(startTime)
		ctrl.logger.Info(
			"Incoming request",
			"client_ip", clientIP,
			"method", method,
			"url", url,
			"duration", duration.String(),
		)
	})
}
