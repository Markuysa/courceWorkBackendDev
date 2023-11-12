package middleware

import "net/http"

type Middleware interface {
	AdminAuth(next http.Handler) http.Handler
	ClientAuth(next http.Handler) http.Handler
}
