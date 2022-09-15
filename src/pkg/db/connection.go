package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	conn *gorm.DB
}

func createConnection(cnfg *Config) *Connection {
	dsn := Resolve(cnfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &Connection{
		conn: db,
	}
}

func Load() {
	cnfgInstance := Config{}
	cnfg := *cnfgInstance.bootCnfg()

	createConnection(&cnfg)
}