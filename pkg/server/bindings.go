package server

import (
	"github.com/Pcorner0/calculadora-avanzada/pkg/api/aritmetica"
	"github.com/Pcorner0/calculadora-avanzada/pkg/api/geometria"
	"gorm.io/gorm"
)

func bindAritmetica(dbHandler *gorm.DB) aritmetica.Handler {
	repository := aritmetica.NewRepository(dbHandler)
	service := aritmetica.NewService(repository)
	return aritmetica.NewHandler(service)
}

func bindGeometria(dbHandler *gorm.DB) geometria.Handler {
	repository := geometria.NewRepository(dbHandler)
	service := geometria.NewService(repository)
	return geometria.NewHandler(service)
}
