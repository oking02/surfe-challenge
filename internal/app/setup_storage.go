package app

import (
	"encoding/json"
	"fmt"
	"os"

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

	var users []domain.User
	if err := json.Unmarshal(blob, &users); err != nil {
		return nil, err
	}
	return users, nil
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

	var actions []domain.Action
	if err := json.Unmarshal(blob, &actions); err != nil {
		return nil, err
	}
	return actions, nil
}
