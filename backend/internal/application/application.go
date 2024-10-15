package application

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/SYSU-ECNC/ecnc-oa/internal/config"
	"github.com/SYSU-ECNC/ecnc-oa/internal/controller"
)

type Application struct {
	cfg    *config.Config
	srv    *http.Server
	logger *slog.Logger
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) InitApplication() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app.logger = logger

	cfg := config.NewConfig(logger)
	cfg.LoadConfig()
	app.cfg = cfg

	ctrl := controller.NewController(logger)
	ctrl.RegisterRoutes()
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerPort),
		Handler:      ctrl.Mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  time.Minute,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}
	app.srv = srv
}

func (app *Application) Run() {
	app.logger.Info("The server is listening", "PORT", app.cfg.ServerPort)
	if err := app.srv.ListenAndServe(); err != nil {
		app.logger.Error("Failed to start the server", "error", err)
	}
}
