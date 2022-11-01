package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
}

func Setup(dbHandler *gorm.DB) {

	router := gin.Default()

	aritmeticHandler := bindAritmetica(dbHandler)

	router.POST("/sumnumbers", aritmeticHandler.SumNumbers)

	router.Run(viper.GetString("server.port"))
}
