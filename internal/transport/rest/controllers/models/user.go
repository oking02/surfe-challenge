package models

import (
	"time"

	"github.com/oking02/surfe-challenge/internal/domain"
)

type User struct {
	ID        domain.UserID `json:"id"`
	Name      string        `json:"name"`
	CreatedAt time.Time     `json:"created_at"`
}

func FromDomainUser(user domain.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}
