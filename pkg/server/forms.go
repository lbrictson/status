package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lbrictson/status/pkg"

	"github.com/lbrictson/status/pkg/auth"

	"github.com/labstack/echo/v4"
)

func (s Server) loginForm(c echo.Context) error {
	type LoginFormValues struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	data := LoginFormValues{}
	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if !authenticateOperator(s.store, data.Email, data.Password) {
		return c.String(http.StatusUnauthorized, "invalid email or password")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Logged in %v successfully", data.Email))
}

func authenticateOperator(s pkg.StorageBackend, email string, password string) bool {
	operator, err := s.GetOperatorByEmail(context.Background(), email)
	if err != nil {
		return false
	}
	if !auth.ComparePassword(password, operator.HashedPassword) {
		return false
	}
	return true
}
