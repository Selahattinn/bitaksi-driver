package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// logMiddleware handles logging
func (a *API) logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"host":       r.Host,
			"address":    r.RemoteAddr,
			"method":     r.Method,
			"requestURI": r.RequestURI,
			"proto":      r.Proto,
			"useragent":  r.UserAgent(),
		}).Info("HTTP request information")

		next.ServeHTTP(w, r)
	})
}

// corsMiddleware handles preflight
func (a *API) corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, HEAD")

		next.ServeHTTP(w, r)
	})
}

// authMiddleware handles authentication
func (a *API) authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get Auth Header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		//Check if its coming from the matchService
		if authHeader == a.config.MatchServiceFlag {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)

	})
}