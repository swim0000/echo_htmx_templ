package repositories

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	db "github.com/swim0000/echo_htmx_templ/db"
	"github.com/swim0000/echo_htmx_templ/models"
)

func InsertSignup(email string, password string) (*models.Signup, error) {
	dbConn := db.ConnectDB()
	defer dbConn.Close()

	query := sq.Insert("signup").
		Columns("email", "password").
		Values(email, password)

	result, err := query.RunWith(dbConn).Exec()
	if err != nil {
		return nil, fmt.Errorf("failed to insert signup: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get LastInsertId: %w", err)
	}
	return GetSignup(id)
}
