package usecase

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/yudai2929/kendai-navi/backend/pkg/domain/entity"
	"github.com/yudai2929/kendai-navi/backend/pkg/domain/repository"
	"github.com/yudai2929/kendai-navi/backend/pkg/lib/auth"
	"github.com/yudai2929/kendai-navi/backend/pkg/lib/errors"
)

type AuthUsecase interface {
	RegisterUser(ctx context.Context, input *RegisterUserInput) error
}

type AuthUsecaseImpl struct {
	authClient     auth.AuthClient
	userRepository repository.UserRepository
	validate       *validator.Validate
	uuid           func() string
	time           func() time.Time
}

func NewAuthUsecase(userRepository repository.UserRepository, authClient auth.AuthClient, validate *validator.Validate) AuthUsecase {
	return &AuthUsecaseImpl{
		authClient:     authClient,
		userRepository: userRepository,
		uuid:           uuid.NewString,
		time:           time.Now,
		validate:       validate,
	}
}

type RegisterUserInput struct {
	Email           string `validate:"required,email"`
	Nickname        string `validate:"required"`
	ProfileImageURL string `validate:"required,url"`
}

func (u *AuthUsecaseImpl) RegisterUser(ctx context.Context, input *RegisterUserInput) error {
	if err := u.validate.Struct(input); err != nil {
		return errors.Wrap(err)
	}

	authUser, err := u.authClient.GetUserByEmail(ctx, input.Email)
	if err != nil && !errors.EqualCode(err, errors.NotFound) {
		return errors.Wrap(err)
	}
	if authUser != nil {
		return errors.Newf(errors.AlreadyExists, "user already exists")
	}

	user, err := u.userRepository.GetUserByEmail(ctx, input.Email)
	if err != nil && !errors.EqualCode(err, errors.NotFound) {
		return errors.Wrap(err)
	}
	if user != nil {
		return errors.Newf(errors.AlreadyExists, "user already exists")
	}

	newAuthUser, err := u.authClient.CreateUser(ctx, input.Email, "password")
	if err != nil {
		return errors.Wrap(err)
	}

	newUser := entity.NewUser(
		u.uuid(),
		input.Nickname,
		newAuthUser.UID,
		input.Email,
		input.ProfileImageURL,
		u.time(),
		u.time(),
	)

	if _, err := u.userRepository.CreateUser(ctx, newUser); err != nil {
		return errors.Wrap(err)
	}

	return nil
}
