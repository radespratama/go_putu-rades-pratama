package users

import (
	"net/http"
	"restful-api/src/database"
	"restful-api/src/database/factories"
	"restful-api/src/models/user"
	"restful-api/src/utils"

	"github.com/labstack/echo/v4"
)

var db = database.Connect()
var factory = factories.UserFactory(db)

func GetAllUsers(echoContext echo.Context) error {
	users, err := factory.GetAllUsers()

	if err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed", err)
	}

	return utils.CreateResponse(echoContext, http.StatusOK, "OK", users)
}

func GetUser(echoContext echo.Context) error {
	var userID string = echoContext.Param("id")
	userData, err := factory.GetUser(userID)

	if err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed", err)
	}

	return utils.CreateResponse(echoContext, http.StatusOK, "OK", userData)
}

func CreateUser(echoContext echo.Context) error {
	var request user.User
	echoContext.Bind(&request)

	if err := factory.CreateUser(request); err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed")
	}

	return utils.CreateResponse(echoContext, http.StatusCreated, "user created")
}

func UpdateUser(echoContext echo.Context) error {
	var request user.User
	echoContext.Bind(&request)

	if err := factory.UpdateUser(request, echoContext.Param("id")); err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed", err)
	}

	return utils.CreateResponse(echoContext, http.StatusNoContent, "user updated")
}

func DeleteUser(echoContext echo.Context) error {
	var userID string = echoContext.Param("id")

	if err := factory.DeleteUser(userID); err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed")
	}

	return utils.CreateResponse(echoContext, http.StatusNoContent, "user deleted")
}
