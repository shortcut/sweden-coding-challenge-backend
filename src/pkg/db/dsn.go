package db

import (
	"fmt"
	"os"
)

type Config struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
	TimeZone     string
}

func (cnfg *Config) bootCnfg() *Config {
	return &Config{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		User:         os.Getenv("DB_USERNAME"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		TimeZone:     os.Getenv("DB_TIMEZONE"),
	}
}

func Resolve(cnfg *Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s sslmode=disable", 
		cnfg.Host, 
		cnfg.User,
		cnfg.Password,
		cnfg.DatabaseName,
		cnfg.Port,
		cnfg.TimeZone,
	)
}