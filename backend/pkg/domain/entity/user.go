package entity

import "time"

type User struct {
	ID              string
	Nickname        string
	FirebaseUID     string
	Email           string
	ProfileImageURL string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewUser(id, nickname, firebaseUID, email, profileImageURL string, createdAt, updatedAt time.Time) *User {
	return &User{
		ID:              id,
		Nickname:        nickname,
		FirebaseUID:     firebaseUID,
		Email:           email,
		ProfileImageURL: profileImageURL,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}
