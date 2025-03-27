package datasources

import (
	"context"

	"github.com/oking02/surfe-challenge/internal/domain"
)

type ActionRepository interface {
	ActionLister
	UserActionLister
}

type UserActionLister interface {
	ListUserActions(context.Context, domain.ClientID, domain.UserID) ([]domain.Action, error)
}

type ActionLister interface {
	ListActions(context.Context, domain.ClientID) ([]domain.Action, error)
}

type ActionProbability interface {
	NextActionProbability(ctx context.Context, clientID domain.ClientID, actionType domain.ActionType) (map[domain.ActionType]float64, error)
}

type ReferralStatistics interface {
	ReferralIndex(ctx context.Context, clientID domain.ClientID) (map[domain.UserID]int, error)
}
