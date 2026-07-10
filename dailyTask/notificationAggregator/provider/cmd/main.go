package main

import (
	"context"
	"log"
	"os/signal"
	"provider/internal/app"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	application := app.New()
	if err := application.Run(ctx); err != nil {
		log.Fatal("Ошибка сервера: ", err)
	}
}
