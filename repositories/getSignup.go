package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	database "github.com/swim0000/echo_htmx_templ/db"
	"github.com/swim0000/echo_htmx_templ/models"
)

func GetSignup(userID int64) (*models.Signup, error) {
	dbConn := database.ConnectDB()
	defer dbConn.Close()

	// Assuming you have a 'users' table
	query := squirrel.Select("*").From("users").Where(squirrel.Eq{"id": userID})

	var signup models.Signup
	err := query.RunWith(dbConn).QueryRow().Scan(&signup.ID, &signup.Email, &signup.Password)

	if err == sql.ErrNoRows {
		return nil, nil // No user found
	} else if err != nil {
		return nil, fmt.Errorf("failed to get signup: %w", err)
	}

	return &signup, nil
}
