package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/oking02/surfe-challenge/internal/app"
)

func main() {
	fmt.Println("Hello World")
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	a, err := app.Setup(ctx)
	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		<-time.After(5 * time.Second)
	}()

	if err = a.Run(ctx); err != nil {
		panic(err)
	}

	os.Exit(1)
}
