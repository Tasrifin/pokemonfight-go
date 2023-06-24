package constants

import "github.com/Tasrifin/pokemonfight-go/helpers"

var (
	APP_PORT = helpers.GetENV("APP_PORT")
	DB_HOST  = helpers.GetENV("DB_HOST")
	DB_NAME  = helpers.GetENV("DB_NAME")
	DB_PASS  = helpers.GetENV("DB_PASS")
	DB_USER  = helpers.GetENV("DB_USER")
	DB_PORT  = helpers.GetENV("DB_PORT")
)
