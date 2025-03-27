package domain

import (
	"iter"
	"maps"
	"time"
)

type ActionID int64

type ActionType string

const (
	ADD_CONTACT   ActionType = "ADD_CONTACT"
	CONNECT_CRM   ActionType = "CONNECT_CRM"
	EDIT_CONTACT  ActionType = "EDIT_CONTACT"
	VIEW_CONTACTS ActionType = "VIEW_CONTACTS"
	REFER_USER    ActionType = "REFER_USER"
	WELCOME       ActionType = "WELCOME"
)

type Action struct {
	// ID is the unique identifier for this action
	ID ActionID
	// Type is a enum value of
	Type ActionType
	// UserID refers to the user who took this action
	UserID UserID
	// TargetID is only for some action types
	// is refers to what user this action is performed on
	TargetID UserID
	// CreatedAt refers when this action was performed
	CreatedAt time.Time
	// ClientID is the identifier for the client
	// the parent user is associated with
	ClientID ClientID
}

func DistinctActionsType(actions []Action) iter.Seq[ActionType] {
	types := make(map[ActionType]struct{})
	for _, action := range actions {
		types[action.Type] = struct{}{}
	}
	return maps.Keys(types)
}
