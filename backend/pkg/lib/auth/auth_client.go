package auth

import "context"

type AuthClient interface {
	GetUserByEmail(ctx context.Context, email string) (*AuthUser, error)
	CreateUser(ctx context.Context, email, password string) (*AuthUser, error)
	DeleteUser(ctx context.Context, uid string) error
}
