package migrations

import (
	"github.com/gustavorocha11/users-crud/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
