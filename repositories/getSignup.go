package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	database "github.com/swim0000/echo_htmx_templ/db"
	"github.com/swim0000/echo_htmx_templ/models"
)

func GetSignup(signupID int64) (*models.Signup, error) {
	dbConn := database.ConnectDB()
	defer dbConn.Close()

	query := squirrel.Select("*").From("signup").Where(squirrel.Eq{"id": signupID})

	var signup models.Signup
	err := query.RunWith(dbConn).QueryRow().Scan(&signup.ID, &signup.Email, &signup.Password)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get signup: %w", err)
	}

	return &signup, nil
}
