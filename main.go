package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/JosiahEdington/gym-log/app"
	"github.com/JosiahEdington/gym-log/app/handler"
	"github.com/JosiahEdington/gym-log/data"
	"github.com/JosiahEdington/gym-log/logs"
)

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	config := app.LoadConfig()
	logger := logs.NewLogger(ctx)
	err := data.ConnectToDB(&config)
	if err != nil {
		return err
	}
	server := handler.NewServer(&config, logger)

	err = handler.StartServer(ctx, server)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	err := run(ctx)
	if err != nil {
		log.Fatalf("Failed with error: %v\n%v", err, ctx)
	}
}
