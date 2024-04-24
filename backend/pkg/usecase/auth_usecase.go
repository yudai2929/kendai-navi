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
	RegisterUser(ctx context.Context, input *RegisterUserInput) (*RegisterUserOutput, error)
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
	Email           string
	Password        string
	Nickname        string
	ProfileImageURL string
}

type RegisterUserOutput struct {
	User  *entity.User
	Token string
}

func (u *AuthUsecaseImpl) RegisterUser(ctx context.Context, input *RegisterUserInput) (output *RegisterUserOutput, err error) {
	if err := u.validate.Struct(input); err != nil {
		return nil, errors.Wrap(err)
	}

	authUser, err := u.authClient.GetUserByEmail(ctx, input.Email)
	if err != nil && !errors.EqualCode(err, errors.NotFound) {
		return nil, errors.Wrap(err)
	}
	if authUser != nil {
		return nil, errors.Newf(errors.AlreadyExists, "user already exists")
	}

	user, err := u.userRepository.GetUserByEmail(ctx, input.Email)
	if err != nil && !errors.EqualCode(err, errors.NotFound) {
		return nil, errors.Wrap(err)
	}
	if user != nil {
		return nil, errors.Newf(errors.AlreadyExists, "user already exists")
	}

	newAuthUser, err := u.authClient.CreateUser(ctx, input.Email, input.Password)
	if err != nil {
		return nil, errors.Wrap(err)
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
	_, createErr := u.userRepository.CreateUser(ctx, newUser)
	if createErr != nil {
		return nil, errors.Wrap(createErr)
	}

	// dbに保存が失敗した場合、authのユーザーも削除する
	defer func() {
		if createErr != nil {
			if err := u.authClient.DeleteUser(ctx, newAuthUser.UID); err != nil {
				// log error
			}
		}
	}()

	return &RegisterUserOutput{
		User:  newUser,
		Token: "token",
	}, nil
}
