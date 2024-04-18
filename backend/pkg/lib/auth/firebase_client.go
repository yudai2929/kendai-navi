package auth

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/yudai2929/kendai-navi/backend/pkg/lib/errors"
)

type FirebaseClient struct {
	*auth.Client
}

func NewFirebaseClient(ctx context.Context) (AuthClient, error) {
	app, err := firebase.NewApp(ctx, nil)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to create firebase app")
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create firebase auth client")
	}
	return &FirebaseClient{client}, nil
}

func (c *FirebaseClient) GetUserByEmail(ctx context.Context, email string) (*AuthUser, error) {
	u, err := c.Client.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return NewAuthUser(u), nil
}

func (c *FirebaseClient) CreateUser(ctx context.Context, email, password string) (*AuthUser, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(password).
		Disabled(false)
	u, err := c.Client.CreateUser(ctx, params)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return NewAuthUser(u), nil
}
