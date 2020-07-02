package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mhdiiilham/gournal/db"
	"github.com/mhdiiilham/gournal/models"
	"github.com/mhdiiilham/gournal/routers"
	"github.com/stretchr/testify/assert"
)

type signupRes struct {
	AccessToken   string `json:"access_token"`
	AdminFullname string `json:"admin_fullname"`
	Message       string `json:"message"`
}

func TestSignUp(t *testing.T) {
	var user models.Admin
	var response signupRes
	var typeCheck signupRes
	body := []byte(`{"fullname": "Testing", "email": "testing@mail.com", "password": "123456"}`)
	r := routers.Router()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "User Created!", response.Message)
	assert.Equal(t, "Testing", response.AdminFullname)
	assert.IsType(t, response.AccessToken, typeCheck.AccessToken)
	assert.IsType(t, response.AdminFullname, typeCheck.AdminFullname)
	assert.IsType(t, response.Message, typeCheck.Message)

	db.DB().Where("email = ?", "testing@mail.com").Delete(&user)
	fmt.Println("RECORD DELETED!")
}
