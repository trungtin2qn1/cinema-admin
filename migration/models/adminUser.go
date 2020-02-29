package models

import (
	"cinema-admin/models"
	"cinema-admin/utils"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gormigrate.v1"
)

//InitAdmin ...
var InitAdmin = &gormigrate.Migration{
	ID: "init_admin",
	Migrate: func(tx *gorm.DB) error {
		var err error

		type adminUser struct {
			ID        uint64 `json:"id,omitempty"`
			Email     string `gorm:"not null;unique"`
			FirstName string
			LastName  string
			Password  string
			Username  string `json:"username,omitempty"`
			Role      string `json:"role,omitempty"`
			LastLogin *time.Time
			CreatedAt time.Time
			UpdatedAt time.Time
			DeletedAt *time.Time
		}

		if err = tx.CreateTable(&adminUser{}).Error; err != nil {
			go utils.LogErrToFile(err.Error())
			return err
		}
		var pwd []byte
		if pwd, err = bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost); err != nil {
			go utils.LogErrToFile(err.Error())
			return err
		}

		usr := adminUser{
			ID:        1,
			Email:     "root@gmail.com",
			FirstName: "root",
			LastName:  "root",
			Username:  "root",
			Role:      models.Role.ROOT,
			Password:  string(pwd),
		}
		return tx.Save(&usr).Error
	},
	Rollback: func(tx *gorm.DB) error {
		go utils.LogErrToFile(err.Error())
		return tx.DropTable("admin_users").Error
	},
}
