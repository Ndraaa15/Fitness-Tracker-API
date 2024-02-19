package server

import (
	"fmt"
	authhandler "github/Ndraaa15/fitness-tracker-api/internal/api/auth/handler/http"
	authstore "github/Ndraaa15/fitness-tracker-api/internal/api/auth/repository/postgresql"
	authservice "github/Ndraaa15/fitness-tracker-api/internal/api/auth/service"
	"github/Ndraaa15/fitness-tracker-api/pkg/config"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Bootstrap struct {
	v   *viper.Viper
	srv *http.Server
	r   *mux.Router
	h   []handler
}

type handler interface {
	Start(mx *mux.Router)
}

func Run() int {
	srv, err := newServer()
	if err != nil {
		log.Printf("Error initializing server: %v", err)
		return -1
	}

	return srv.Start()

}

func newServer() (*Bootstrap, error) {
	srv := &Bootstrap{
		srv: &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}

	viper, err := config.InitializeViper()
	if err != nil {
		log.Printf("Error initializing viper: %v", err)
		return nil, fmt.Errorf("Error initializing viper: %v", err)
	}

	srv.v = viper

	db, err := config.InitializeDatabaseConn(viper)
	if err != nil {
		log.Printf("Error initializing database: %v", err)
		return nil, fmt.Errorf("Error initializing database: %v", err)
	}

	srv.r = mux.NewRouter()

	srv.RegisterRoutes(db)

	return srv, nil
}

func (s *Bootstrap) RegisterRoutes(db *sqlx.DB) error {
	var authService authservice.AuthServiceImpl
	{
		authStore := authstore.NewStore(db)
		authService = authservice.NewAuthService(authStore)
	}

	{
		authIdentities := []authhandler.HandlerIdetifier{
			authhandler.HandlerRegister,
			authhandler.HandlerLogin,
		}

		authHTTPHandler, err := authhandler.New(authService, authIdentities...)
		if err != nil {
			log.Printf("Error initializing auth handler: %v", err)
			return fmt.Errorf("Error initializing auth handler: %v", err)
		}

		s.h = append(s.h, authHTTPHandler)
	}

	return nil
}

func (s *Bootstrap) Start() int {
	appMux := s.r.PathPrefix("/api/v1").Subrouter()
	appMux.Use(corsMiddleware)

	for _, handler := range s.h {
		handler.Start(appMux)
	}

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", s.v.GetString("ADDRESS"), s.v.GetString("PORT")), s.r)
	if err != nil {
		log.Printf("Error starting server: %v", err)
		return -1
	}

	return 0
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
