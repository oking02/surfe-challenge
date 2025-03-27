package memory

import (
	"context"

	"github.com/oking02/surfe-challenge/internal/domain"
)

type ActionRepository struct {
	userActions map[userKey][]domain.Action
}

func NewActionRepository(data []domain.Action) *ActionRepository {

	userActions := make(map[userKey][]domain.Action)
	for _, action := range data {

		key := userKey{
			clientID: action.ClientID,
			userID:   action.UserID,
		}

		if actions, ok := userActions[key]; ok {
			userActions[key] = append(actions, action)
			continue
		}
		userActions[key] = []domain.Action{action}
	}

	return &ActionRepository{
		userActions: userActions,
	}
}

func (ar *ActionRepository) ListUserActions(ctx context.Context, clientID domain.ClientID, user domain.UserID) ([]domain.Action, error) {
	var results []domain.Action
	actions, ok := ar.userActions[userKey{
		clientID: clientID,
		userID:   user,
	}]
	if ok {
		results = actions
	}
	return results, nil
}

func (ar *ActionRepository) ListActions(ctx context.Context, clientID domain.ClientID) ([]domain.Action, error) {
	var results []domain.Action
	for key, actions := range ar.userActions {
		if key.clientID == clientID {
			results = append(results, actions...)
		}
	}
	return results, nil
}
