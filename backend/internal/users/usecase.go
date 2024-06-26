package users

import (
	"context"
	"runner-manager-backend/internal/users/dto"
)

type Usecase interface {
	Login(ctx context.Context, request *dto.UserLoginRequest) (response *dto.UserLoginResponse, err error)
	LoginViaApiKey(ctx context.Context, request *dto.UserLoginApiKeyRequest) (response *dto.UserLoginResponse, err error)
	Create(ctx context.Context, payload *dto.CreateUserRequest) (userID string, err error)
	GenerateApiKey(ctx context.Context, userID string) (apiKey string, err error)
}
