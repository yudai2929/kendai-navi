package psql

import (
	"context"

	"github.com/yudai2929/kendai-navi/backend/db/ent"
	"github.com/yudai2929/kendai-navi/backend/pkg/domain/entity"
	"github.com/yudai2929/kendai-navi/backend/pkg/domain/repository"
	db "github.com/yudai2929/kendai-navi/backend/pkg/lib/ent"
	"github.com/yudai2929/kendai-navi/backend/pkg/lib/errors"
)

type UserRepositoryImpl struct {
	client *db.Client
}

func NewUserRepositoryImpl(client *db.Client) repository.UserRepository {
	return &UserRepositoryImpl{client: client}
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	u, err := r.client.User.Create().
		SetID(user.ID).
		SetEmail(user.Email).
		SetNickname(user.Nickname).
		SetProfileImageURL(user.ProfileImageURL).
		SetCreatedAt(user.CreatedAt).
		SetUpdatedAt(user.UpdatedAt).
		Save(ctx)

	if err != nil {
		return nil, errors.Wrap(err)
	}
	return toEntity(u), nil
}

func toEntity(u *ent.User) *entity.User {
	return &entity.User{
		ID:              u.ID,
		Email:           u.Email,
		Nickname:        u.Nickname,
		ProfileImageURL: u.ProfileImageURL,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
	}
}
