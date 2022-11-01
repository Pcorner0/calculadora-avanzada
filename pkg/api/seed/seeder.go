package seed

import (
	"github.com/Pcorner0/calculadora-avanzada/pkg/models"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Users{})
	if err != nil {
		log.Fatal(err)
		return
	}
}
