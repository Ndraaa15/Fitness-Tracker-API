package service

import "github/Ndraaa15/fitness-tracker-api/internal/api/auth/repository"

type AuthService struct {
	authStore repository.AuthStoreImpl
}

func NewAuthService(newAuthStore repository.AuthStoreImpl) AuthServiceImpl {
	return &AuthService{
		authStore: newAuthStore,
	}
}

type AuthServiceImpl interface {
}
