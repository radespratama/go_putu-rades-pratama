package factories

import (
	"restful-api/src/models/auth"

	"gorm.io/gorm"
)

func LoginFactory(db *gorm.DB) auth.LoginModel {
	return &auth.LoginRepository{
		DB: db,
	}
}
