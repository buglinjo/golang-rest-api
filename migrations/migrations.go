package migrations

import (
	"github.com/buglinjo/golang-rest-api/app/models"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
