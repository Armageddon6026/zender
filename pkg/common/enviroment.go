package common

import "os"

var (
	JWT_SECRET  string = "zender6026"
	DB_USER     string
	DB_PASSWORD string
	DB_ADDR     string
)

func init() {
	if theJwtSecret := os.Getenv("JWT_SECRET"); theJwtSecret != "" {
		JWT_SECRET = theJwtSecret
	}
	if addr := os.Getenv("DB_ADDR"); addr != "" {
		DB_USER = addr
	}
	if password := os.Getenv("DB_PASSWORD"); password != "" {
		DB_PASSWORD = password
	}
	if user := os.Getenv("DB_USER"); user != "" {
		DB_USER = user
	}
}
