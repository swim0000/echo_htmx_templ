package repositories

import (
	"log"

	database "github.com/swim0000/echo_htmx_templ/db"
	"github.com/swim0000/echo_htmx_templ/models"
)

func InsertSignup(signup *models.Signup) error {
	dbConn := database.GetDB()
	defer dbConn.Close()

	stmt, err := dbConn.Prepare("INSERT INTO signup (user_email, user_password) VALUES ($1, $2)")
	if err != nil {
		log.Panicln(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(signup.Email, signup.Password)
	if err != nil {
		log.Panicln(err)
		return err
	}

	log.Println("User registered successfully!")

	return nil
}
