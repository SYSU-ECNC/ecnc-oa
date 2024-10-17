package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/SYSU-ECNC/ecnc-oa/backend/internal/config"
	"github.com/SYSU-ECNC/ecnc-oa/backend/internal/controller"
	"github.com/SYSU-ECNC/ecnc-oa/backend/internal/repository"
	"github.com/SYSU-ECNC/ecnc-oa/backend/internal/service"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg := config.NewConfig(logger)
	cfg.LoadConfig()

	dsn := fmt.Sprintf("postgresql://postgres:password@localhost:%s/ecnc_oa_db?sslmode=disable", cfg.DatabasePassword)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Error("Cannot open database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	ctrl := controller.NewController(logger, svc)
	ctrl.RegisterRoutes()
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerPort),
		Handler:      ctrl.Handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  time.Minute,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("The server is listening", "PORT", cfg.ServerPort)
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("Failed to start the server", "error", err)
	}
}
