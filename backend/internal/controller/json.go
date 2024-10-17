package controller

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (ctrl *Controller) readJSON(r *http.Request, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
}

func (ctrl *Controller) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		ctrl.logger.Error("Some error occured in writing JSON", "error", err)
		// don't use internalServerError function to avoid infinite loop
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (ctrl *Controller) writeSuccessJSON(w http.ResponseWriter, status int, message string, data any) {
	type envelope struct {
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}
	ctrl.writeJSON(w, status, envelope{
		Message: message,
		Data:    data,
	})
}

func (ctrl *Controller) writeErrorJSON(w http.ResponseWriter, status int, err error) {
	type envelope struct {
		Error string `json:"error"`
	}
	ctrl.writeJSON(w, status, envelope{
		Error: err.Error(),
	})
}

func (ctrl *Controller) internalServerError(w http.ResponseWriter) {
	ctrl.writeErrorJSON(w, http.StatusInternalServerError, errors.New("服务器内部错误"))
}
