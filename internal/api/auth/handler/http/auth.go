package http

import (
	"github/Ndraaa15/fitness-tracker-api/internal/api/auth/service"
	"net/http"
)

type authHandler struct {
	authService service.AuthServiceImpl
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/auth/signup":
		h.Signup(w, r)
	case "/auth/signin":
		h.Signin(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *authHandler) Signup(w http.ResponseWriter, r *http.Request) {

}

func (h *authHandler) Signin(w http.ResponseWriter, r *http.Request) {

}
