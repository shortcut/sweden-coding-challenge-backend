package env

import "github.com/joho/godotenv"

func Load() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
}