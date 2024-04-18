package auth

import "firebase.google.com/go/v4/auth"

type AuthUser struct {
	UID           string
	Email         string
	EmailVerified bool
}

func NewAuthUser(u *auth.UserRecord) *AuthUser {
	return &AuthUser{
		UID:           u.UID,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
	}
}
