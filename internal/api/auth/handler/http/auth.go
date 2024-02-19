package http

import (
	"github/Ndraaa15/fitness-tracker-api/internal/api/auth/service"
	"net/http"
)

type authHandler struct {
	authService service.AuthServiceImpl
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
