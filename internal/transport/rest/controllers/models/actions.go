package models

import (
	"time"

	"github.com/oking02/surfe-challenge/internal/domain"
)

type Action struct {
	ID         domain.UserID `json:"id"`
	Type       string        `json:"name"`
	CreatedAt  time.Time     `json:"created_at"`
	TargetUser domain.UserID `json:"target_user,omitempty"`
}

func FromDomainActions(action []domain.Action) []Action {
	results := make([]Action, len(action))
	for i, a := range action {
		results[i] = FromDomainAction(a)
	}
	return results
}

func FromDomainAction(action domain.Action) Action {
	return Action{
		ID:         domain.UserID(action.ID),
		Type:       string(action.Type),
		CreatedAt:  action.CreatedAt,
		TargetUser: action.TargetID,
	}
}
