package server

import (
	"github.com/Pcorner0/calculadora-avanzada/pkg/api/aritmetica"
	"gorm.io/gorm"
)

func bindAritmetica(dbHandler *gorm.DB) aritmetica.Handler {
	repository := aritmetica.NewRepository(dbHandler)
	service := aritmetica.NewService(repository)
	return aritmetica.NewHandler(service)
}
