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
	geometricHandler := bindGeometria(dbHandler)

	router.POST("/sumsubtnumbers", aritmeticHandler.SumSubtNumbers)
	router.POST("/divnumbers", aritmeticHandler.DivideTwoNumbers)
	router.POST("/mulnumbers", aritmeticHandler.MultiplyNumbers)
	router.POST("/root", aritmeticHandler.RootNumbers)
	router.POST("/power", aritmeticHandler.PowerNumbers)
	router.POST("/sin",geometricHandler.SinNumbers)
	router.POST("/cos",geometricHandler.CosNumbers)
	router.POST("/tan",geometricHandler.TanNumbers)
	
	router.Run(viper.GetString("server.port"))
}

