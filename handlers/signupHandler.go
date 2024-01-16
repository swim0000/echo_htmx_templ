package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swim0000/echo_htmx_templ/models"
	"github.com/swim0000/echo_htmx_templ/repositories"
)

func SignupHandler(c echo.Context) error {
	form := new(models.Signup)
	if err := c.Bind(form); err != nil {
		return err
	}

	_, err := repositories.InsertSignup(form.Email, form.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save signup"})
	}

	return c.JSON(http.StatusOK, form)
}
