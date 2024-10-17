package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (ctrl *Controller) login(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req request
	if err := ctrl.readJSON(r, &req); err != nil {
		ctrl.writeErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if req.Username == "" {
		ctrl.writeErrorJSON(w, http.StatusBadRequest, ErrUsernameEmpty)
		return
	}
	if req.Password == "" {
		ctrl.writeErrorJSON(w, http.StatusBadRequest, ErrPasswordEmpty)
		return
	}

	// check password
	user, err := ctrl.repo.GetUserByUsername(r.Context(), req.Username)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			ctrl.writeErrorJSON(w, http.StatusNotFound, ErrUsernameNotExists)
		default:
			ctrl.logger.Error("Unexpect error in login", "error", err)
			ctrl.internalServerError(w)
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			ctrl.writeErrorJSON(w, http.StatusUnauthorized, ErrWrongPassword)
		default:
			ctrl.logger.Error("Unexpect error in login", "error", err)
			ctrl.internalServerError(w)
		}
		return
	}

	// generate jwt
	expires := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   user.Username,
		ExpiresAt: jwt.NewNumericDate(expires),
	})
	ss, err := token.SignedString([]byte(ctrl.cfg.JWTSecret))
	if err != nil {
		ctrl.logger.Error("Unexpect error in login", "error", err)
		ctrl.internalServerError(w)
		return
	}

	// set jwt to http-only cookie
	cookie := &http.Cookie{
		Name:     "__ecnc_oa_jwt",
		Value:    ss,
		Path:     "/",
		Expires:  expires,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	ctrl.writeSuccessJSON(w, http.StatusOK, "登录成功", nil)
}

func (ctrl *Controller) logout(w http.ResponseWriter, r *http.Request) {
	expires := time.Now().Add(-time.Hour)
	cookie := &http.Cookie{
		Name:     "__ecnc_oa_jwt",
		Value:    "",
		Path:     "/",
		Expires:  expires,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	ctrl.writeSuccessJSON(w, http.StatusOK, "登出成功", nil)
}
