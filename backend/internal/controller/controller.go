package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Controller struct {
	logger *slog.Logger
	Mux    *http.ServeMux
}

func NewController(logger *slog.Logger) *Controller {
	return &Controller{logger: logger, Mux: http.NewServeMux()}
}

func (ctrl *Controller) RegisterRoutes() {
	v1Mux := http.NewServeMux()
	v1Mux.Handle("GET /hello-world", ctrl.logging(http.HandlerFunc(ctrl.helloWorld)))

	ctrl.Mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1Mux))
}

func (ctrl *Controller) helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello, world!",
	})
}
