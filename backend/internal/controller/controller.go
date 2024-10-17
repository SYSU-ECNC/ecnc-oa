package controller

import (
	"log/slog"
	"net/http"

	"github.com/SYSU-ECNC/ecnc-oa/backend/internal/config"
	"github.com/SYSU-ECNC/ecnc-oa/backend/internal/repository"
)

type Controller struct {
	cfg    *config.Config
	logger *slog.Logger
	repo   *repository.Repository

	Handler http.Handler
}

func NewController(cfg *config.Config, logger *slog.Logger, repo *repository.Repository) *Controller {
	return &Controller{cfg: cfg, logger: logger, repo: repo}
}

func (ctrl *Controller) RegisterRoutes() {
	v1Mux := http.NewServeMux()
	v1Mux.Handle("POST /auth/login", http.HandlerFunc(ctrl.login))
	v1Mux.Handle("DELETE /auth/logout", http.HandlerFunc(ctrl.logout))

	main := http.NewServeMux()
	main.Handle("/api/v1/", http.StripPrefix("/api/v1", v1Mux))

	ctrl.Handler = ctrl.logging(main)
}
