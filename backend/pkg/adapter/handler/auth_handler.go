package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yudai2929/kendai-navi/backend/pkg/lib/errors"
	"github.com/yudai2929/kendai-navi/backend/pkg/usecase"
)

type AuthHandler interface {
	RegisterUser(c *gin.Context) (*RegisterUserResponse, error)
}

type AuthHandlerImpl struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) AuthHandler {
	return &AuthHandlerImpl{
		authUsecase: authUsecase,
	}
}

type RegisterUserRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	Nickname        string `json:"nickname"`
	ProfileImageURL string `json:"profile_image_url"`
}

type RegisterUserResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type UserResponse struct {
	ID              string `json:"id"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	ProfileImageURL string `json:"profile_image_url"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func (h *AuthHandlerImpl) RegisterUser(c *gin.Context) (*RegisterUserResponse, error) {
	req := RegisterUserRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, errors.Wrap(err)
	}

	ctx := c.Request.Context()
	input := usecase.RegisterUserInput{
		Email:           req.Email,
		Password:        req.Password,
		Nickname:        req.Nickname,
		ProfileImageURL: req.ProfileImageURL,
	}

	output, err := h.authUsecase.RegisterUser(ctx, &input)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return h.toRegisterUserResponse(output), nil
}

func (h *AuthHandlerImpl) toRegisterUserResponse(output *usecase.RegisterUserOutput) *RegisterUserResponse {
	return &RegisterUserResponse{
		User: UserResponse{
			ID:              output.User.ID,
			Nickname:        output.User.Nickname,
			Email:           output.User.Email,
			ProfileImageURL: output.User.ProfileImageURL,
			CreatedAt:       output.User.CreatedAt.String(),
			UpdatedAt:       output.User.UpdatedAt.String(),
		},
		Token: output.Token,
	}
}
