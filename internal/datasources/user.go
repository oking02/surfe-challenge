package datasources

import (
	"context"

	"github.com/oking02/surfe-challenge/internal/domain"
)

type UserRepository interface {
	UserLister
	UserFetcher
}

type UserLister interface {
	ListUsers(context.Context, domain.ClientID) ([]domain.User, error)
}

type UserFetcher interface {
	GetUser(context.Context, domain.ClientID, domain.UserID) (domain.User, error)
}
