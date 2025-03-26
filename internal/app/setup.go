package app

import (
	"context"
	"net/http"

	"github.com/oking02/surfe-challenge/internal/transport/rest/controllers"
)

func Setup(ctx context.Context) (*App, error) {

	restServer := setupRestServer()

	userRepo, err := setupUserStorage()
	if err != nil {
		return nil, err
	}

	restServer.SetupRoutes(map[string]http.Handler{
		"GET /api/v1/users/{id}": controllers.NewGetUserController(userRepo),
	})

	return &App{
		components: []Component{
			restServer,
		},
	}, nil

}
