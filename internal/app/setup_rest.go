package app

import (
	"github.com/oking02/surfe-challenge/internal/enviroment"
	"github.com/oking02/surfe-challenge/internal/transport/rest"
)

func setupRestServer() *rest.Server {

	return rest.NewServer(
		enviroment.Int("HTTP_PORT", 3000),
	)

}
