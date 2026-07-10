package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"provider/internal/http/handler"
	"provider/internal/server"
	"provider/internal/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type App struct {
	srv *http.Server
}

func New() *App {
	validate := validator.New()
	emailService := service.NewEmailService()
	handlerApp := handler.NewHandler(validate, emailService)
	routerApp := handler.NewRouter(handlerApp)
	serverApp := server.NewServer(routerApp)

	return &App{
		srv: serverApp,
	}
}

func (app *App) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		if err := app.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	log.Println("Сервер запущен на :8080")

	select {
	case <-ctx.Done():
		log.Println("Получен сигнал завершения")

		return app.shutdown(context.Background())

	case err := <-errCh:
		return err
	}
}

func (app *App) shutdown(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := app.srv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	// Здесь также закрываем соединения с БД и фоновые задачи (например, db.Close())

	log.Println("Сервер успешно остановлен")

	return nil
}
