package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetUserController(c echo.Context) error {
	for _, u := range users {
		if strconv.Itoa(u.Id) == c.Param("id") {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user",
				"user":     u,
			})
		}
	}
	return c.JSON(http.StatusBadRequest, "ID User doesn't exist.")
}

func DeleteUserController(c echo.Context) error {
	for i, u := range users {
		if strconv.Itoa(u.Id) == c.Param("id") {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete user",
			})
		}
	}
	return c.JSON(http.StatusBadRequest, "ID doesn't exist.")
}

func UpdateUserController(c echo.Context) error {
	data := new(User)
	if err := c.Bind(data); err != nil {
		return err
	}
	for i, u := range users {
		if strconv.Itoa(u.Id) == c.Param("id") {
			if data.Username != "" {
				users[i].Username = data.Username
			}
			if data.Email != "" {
				users[i].Email = data.Password
			}
			if data.Password != "" {
				users[i].Password = data.Password
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update user",
				"user":     users[i],
			})
		}
	}
	return c.JSON(http.StatusBadRequest, "ID doesn't exist.")
}

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}
