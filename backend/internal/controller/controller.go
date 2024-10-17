package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/SYSU-ECNC/ecnc-oa/backend/internal/service"
)

type Controller struct {
	logger *slog.Logger
	svc    *service.Service

	Handler http.Handler
}

func NewController(logger *slog.Logger, svc *service.Service) *Controller {
	return &Controller{logger: logger, svc: svc}
}

func (ctrl *Controller) RegisterRoutes() {
	v1Mux := http.NewServeMux()
	v1Mux.Handle("GET /hello-world", http.HandlerFunc(ctrl.helloWorld))

	main := http.NewServeMux()
	main.Handle("/api/v1/", http.StripPrefix("/api/v1", v1Mux))

	ctrl.Handler = ctrl.logging(main)
}

func (ctrl *Controller) helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello, world!",
	})
}
