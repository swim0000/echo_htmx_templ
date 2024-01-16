package repositories

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	database "github.com/swim0000/echo_htmx_templ/db"
	"github.com/swim0000/echo_htmx_templ/models"
)

func InsertSignup(email string, password string) (*models.Signup, error) {
	dbConn := database.GetDB()
	defer dbConn.Close()

	query := sq.Insert("signup").
		Columns("user_email", "user_password").
		Values(email, password)
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to generate SQL query: %w", err)
	}
	fmt.Println("Generated SQL query:", sql, args)

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
