package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swim0000/echo_htmx_templ/util"
	views "github.com/swim0000/echo_htmx_templ/views/templates"
)

func SignupFormHandler(c echo.Context) error {
	return util.Render(c, http.StatusOK, views.SignupForm())
}
