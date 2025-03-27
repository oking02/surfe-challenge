package app

import (
	"context"
	"net/http"

	"github.com/oking02/surfe-challenge/internal/command"

	"github.com/oking02/surfe-challenge/internal/transport/rest/controllers"
)

func Setup(ctx context.Context) (*App, error) {

	restServer := setupRestServer()

	userRepo, err := setupUserStorage()
	if err != nil {
		return nil, err
	}

	actionsRepo, err := setupActionsStorage()
	if err != nil {
		return nil, err
	}

	nextActionCmd := command.NewNextActionProbabilityCommand(actionsRepo)

	getUserActionsCmd := command.NewGetUserActionsCommand(userRepo, actionsRepo)

	referralIndexCmd := command.NewReferralIndexCommand(userRepo, actionsRepo)

	restServer.SetupRoutes(map[string]http.Handler{
		"GET /api/v1/users/{id}":                        controllers.NewGetUserController(userRepo),
		"GET /api/v1/users/{id}/actions":                controllers.NewGetUserActionsController(getUserActionsCmd),
		"GET /api/v1/actions/{type}/probability/next":   controllers.NewGetNextActionProbabilityController(nextActionCmd),
		"GET /api/v1/actions/REFER_USER/referral-index": controllers.NewGetReferralIndexController(referralIndexCmd),
	})

	return &App{
		components: []Component{
			restServer,
		},
	}, nil

}
