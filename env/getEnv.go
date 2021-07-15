package env

import "github.com/joho/godotenv"

func SetEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
