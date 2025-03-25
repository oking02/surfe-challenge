package datasources

import (
	"context"
	"github.com/oking02/surfe-challenge/internal/domain"
)

type ActionRepository struct {
	UserActionLister
}

type UserActionLister interface {
	ListUserActions(context.Context, domain.ClientID, domain.UserID) ([]domain.Action, error)
}
