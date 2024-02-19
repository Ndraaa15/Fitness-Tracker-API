package http

import (
	"fmt"
	"github/Ndraaa15/fitness-tracker-api/internal/api/auth/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	handlers    map[string]*handler
	authService service.AuthServiceImpl
}

type handler struct {
	h          http.Handler
	identifier HandlerIdetifier
}

type HandlerIdetifier struct {
	Name string
	URL  string
}

var (
	HandlerRegister = HandlerIdetifier{
		Name: "Signup",
		URL:  "/signup",
	}

	HandlerLogin = HandlerIdetifier{
		Name: "Signin",
		URL:  "/signin",
	}
)

func New(service service.AuthServiceImpl, identifiers ...HandlerIdetifier) (*Handler, error) {
	h := &Handler{
		authService: service,
		handlers:    make(map[string]*handler),
	}

	for _, identifier := range identifiers {
		if h.handlers == nil {
			h.handlers = map[string]*handler{}
		}

		h.handlers[identifier.Name] = &handler{
			identifier: identifier,
		}

		handler, err := h.createHandler(identifier.Name)
		if err != nil {
			return nil, err
		}

		h.handlers[identifier.Name].h = handler
	}

	return h, nil
}

func (h *Handler) createHandler(identifier string) (http.Handler, error) {
	var httpHandler http.Handler

	switch identifier {

	case HandlerRegister.Name:
		httpHandler = &authHandler{
			authService: h.authService,
		}

	default:
		return nil, fmt.Errorf("handler not found")
	}

	return httpHandler, nil
}

func (h *Handler) Start(mx *mux.Router) {
	for _, handler := range h.handlers {
		mx.HandleFunc(handler.identifier.URL, handler.h)
	}
}
