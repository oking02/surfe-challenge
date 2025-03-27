package command

import (
	"context"

	"github.com/oking02/surfe-challenge/internal/datasources"
	"github.com/oking02/surfe-challenge/internal/domain"
)

type GetUserActionsCommand struct {
	userFetcher    datasources.UserFetcher
	actionsFetcher datasources.UserActionLister
}

func NewGetUserActionsCommand(
	userFetcher datasources.UserFetcher,
	actionsFetcher datasources.UserActionLister) *GetUserActionsCommand {
	return &GetUserActionsCommand{
		userFetcher:    userFetcher,
		actionsFetcher: actionsFetcher,
	}
}

func (cmd GetUserActionsCommand) ListUserActions(ctx context.Context, clientID domain.ClientID, userID domain.UserID) ([]domain.Action, error) {

	// check user exists
	_, err := cmd.userFetcher.GetUser(ctx, clientID, userID)
	if err != nil {
		return nil, err
	}

	return cmd.actionsFetcher.ListUserActions(ctx, clientID, userID)
}
