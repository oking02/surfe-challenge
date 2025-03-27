package app

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/oking02/surfe-challenge/internal/datasources"
	"github.com/oking02/surfe-challenge/internal/datasources/memory"
	"github.com/oking02/surfe-challenge/internal/domain"
	"github.com/oking02/surfe-challenge/internal/enviroment"
)

func setupUserStorage() (datasources.UserRepository, error) {

	userData, err := setupUserData()
	if err != nil {
		return nil, fmt.Errorf("failed to get user data: %w", err)
	}

	storageDriver := enviroment.String("STORAGE_DRIVER", "memory")

	switch storageDriver {
	case "memory":
		return memory.NewUserRepository(userData), nil
	case "sqlite":
		// as an example
		// add setup logic
		fallthrough
	case "mysql":
		// as an example,
		// add setup logic
		fallthrough
	default:
		return nil, fmt.Errorf("unknown storage driver: %s", storageDriver)
	}

}

func setupUserData() ([]domain.User, error) {

	userDataFilepath := enviroment.String("USER_DATA_LOCATION", "data/users.json")
	if userDataFilepath == "" {
		return nil, nil
	}

	blob, err := os.ReadFile(userDataFilepath)
	if err != nil {
		return nil, err
	}

	var users []user
	if err := json.Unmarshal(blob, &users); err != nil {
		return nil, err
	}

	results := make([]domain.User, len(users))
	for i, a := range users {
		results[i] = a.toDomain()
	}

	return results, nil
}

func setupActionsStorage() (datasources.ActionRepository, error) {

	actionData, err := setupActionsData()
	if err != nil {
		return nil, fmt.Errorf("failed to get action data: %w", err)
	}

	storageDriver := enviroment.String("STORAGE_DRIVER", "memory")

	switch storageDriver {
	case "memory":
		return memory.NewActionRepository(actionData), nil
	case "sqlite":
		// as an example
		// add setup logic
		fallthrough
	case "mysql":
		// as an example,
		// add setup logic
		fallthrough
	default:
		return nil, fmt.Errorf("unknown storage driver: %s", storageDriver)
	}

}

func setupActionsData() ([]domain.Action, error) {

	actionsDataFilepath := enviroment.String("ACTION_DATA_LOCATION", "data/actions.json")
	if actionsDataFilepath == "" {
		return nil, nil
	}

	blob, err := os.ReadFile(actionsDataFilepath)
	if err != nil {
		return nil, err
	}

	var actions []action
	if err := json.Unmarshal(blob, &actions); err != nil {
		return nil, err
	}

	results := make([]domain.Action, len(actions))
	for i, a := range actions {
		results[i] = a.toDomain()
	}

	return results, nil
}

type user struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u user) toDomain() domain.User {
	return domain.User{
		ID:        domain.UserID(u.ID),
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
	}
}

type action struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	UserID   int    `json:"userId"`
	TargetID int    `json:"targetUser"`
	// CreatedAt refers when this action was performed
	CreatedAt time.Time `json:"createdAt"`
	// ClientID is the identifier for the client
	// the parent user is associated with
	ClientID string
}

func (a *action) toDomain() domain.Action {
	return domain.Action{
		ID:        domain.ActionID(a.ID),
		Type:      domain.ActionType(a.Type),
		UserID:    domain.UserID(a.UserID),
		TargetID:  domain.UserID(a.TargetID),
		CreatedAt: a.CreatedAt,
	}
}
