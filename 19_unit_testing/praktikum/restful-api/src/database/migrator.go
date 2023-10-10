package database

import (
	bookModel "restful-api/src/models/book"
	userModel "restful-api/src/models/user"
)

func Migrate() {
	db := Connect()
	db.AutoMigrate(&userModel.User{}, &bookModel.Book{})
}

func Demigrate() {
	db := Connect()
	db.Migrator().DropTable(&userModel.User{}, &bookModel.Book{})
}
