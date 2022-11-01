package main

import (
	"github.com/Pcorner0/calculadora-avanzada/config"
	"github.com/Pcorner0/calculadora-avanzada/db"
	"github.com/Pcorner0/calculadora-avanzada/pkg/api/seed"
	"github.com/Pcorner0/calculadora-avanzada/pkg/server"
	"github.com/spf13/viper"
)

func main() {
	config.Init()
	dbHandler := db.Init()

	if viper.GetBool("database.migrate") {
		seed.Migrate(dbHandler)
	}

	server.Setup(dbHandler)
}
