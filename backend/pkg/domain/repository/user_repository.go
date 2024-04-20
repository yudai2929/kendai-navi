package repository

import (
	"context"

	"github.com/yudai2929/kendai-navi/backend/pkg/domain/entity"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}
