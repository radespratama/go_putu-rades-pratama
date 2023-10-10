package routes

import (
	"restful-api/src/controllers/auth"
	"restful-api/src/controllers/books"
	"restful-api/src/controllers/users"
	"restful-api/src/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()
	middlewares.Logger(route)

	authRoutes := route.Group("")
	authRoutes.Use(middlewares.VerifyAuthentication())

	route.POST("/login", auth.LoginHandler)

	route.GET("/users", users.GetAllUsers)
	authRoutes.POST("/users", users.CreateUser)
	authRoutes.GET("/users/:id", users.GetUser)
	authRoutes.PUT("/users/:id", users.UpdateUser)
	authRoutes.DELETE("/users/:id", users.DeleteUser)

	authRoutes.GET("/books", books.GetAllBooks)
	authRoutes.POST("/books", books.CreateBook)
	authRoutes.GET("/books/:id", books.GetBook)
	authRoutes.PUT("/books/:id", books.UpdateBook)
	authRoutes.DELETE("/books/:id", books.DeleteBook)

	return route
}
