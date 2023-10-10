package auth

import (
	"log"
	"net/http"
	"restful-api/src/database"
	"restful-api/src/database/factories"
	"restful-api/src/models/auth"
	"restful-api/src/utils"

	"github.com/labstack/echo/v4"
)

var db = database.Connect()
var factory = factories.LoginFactory(db)

func LoginHandler(echoContext echo.Context) error {
	var loginData auth.Login
	echoContext.Bind(&loginData)

	userID, err := factory.AttemptLogin(loginData)
	log.Println(userID, err)
	if err != nil || userID == "" {
		log.Println("Unidentified authentication trials", err)
		return utils.CreateResponse(echoContext, http.StatusUnauthorized, "Can't verify authentication trial.")
	}

	token, _ := utils.GenerateJwt(userID)
	utils.SetAuthCookie(echoContext, token)
	return utils.CreateResponse(echoContext, http.StatusOK, "OK", token)
}
