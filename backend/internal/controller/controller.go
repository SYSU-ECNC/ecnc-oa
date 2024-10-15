package controller

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	Mux *http.ServeMux
}

func NewController() *Controller {
	return &Controller{Mux: http.NewServeMux()}
}

func (ctrl *Controller) RegisterRoutes() {
	ctrl.Mux.HandleFunc("GET /hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, world!",
		})
	})
}
