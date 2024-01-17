package handlers

import (
	"fmt"
	"net/http"

	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
	"github.com/swim0000/echo_htmx_templ/models"
	"github.com/swim0000/echo_htmx_templ/repositories"
	"github.com/swim0000/echo_htmx_templ/util"
	views "github.com/swim0000/echo_htmx_templ/views/templates"
)

func SignupHandler(c echo.Context) error {
	if htmx.IsHTMX(c.Request()) {
		form := new(models.Signup)
		if err := c.Bind(form); err != nil {
			fmt.Println("Error binding form data:", err)
			return err
		}

		if err := repositories.InsertSignup(form); err != nil {
			fmt.Println("Error inserting signup:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save signup"})
		}

		return c.JSON(http.StatusOK, form)
	} else {
		return util.Render(c, http.StatusOK, views.SignupForm())
	}
}
