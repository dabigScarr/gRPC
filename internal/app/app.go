package grpc

import (
	"app/internal/app/grpc"
	"app/internal/services/auth"
	"app/internal/storage/sqlite"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *app.App
}

func New(
	log *slog.Logger, grpcPort int,
	storagePath string,
	tokenTTL time.Duration) *App {

	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := app.New(log, grpcPort, authService)

	return &App{
		GRPCServer: grpcApp,
	}
}
