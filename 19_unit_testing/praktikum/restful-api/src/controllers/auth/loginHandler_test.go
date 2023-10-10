package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restful-api/src/database"
	"restful-api/src/database/factories"
	"restful-api/src/models/auth"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var dbCon = database.Connect()
var factoryLogin = factories.LoginFactory(db)

func TestLoginHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	loginData := auth.Login{
		Email:    "testyahuu@example.com",
		Password: "password",
	}
	c.Set("loginData", loginData)

	err := LoginHandler(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "OK", response["message"])
	assert.NotNil(t, response["token"])
}
