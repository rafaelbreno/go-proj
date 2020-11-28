package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-proj/app/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestStruct struct {
	Auth   Auth
	List   List
	Router *gin.Engine
}

type Auth struct {
	Data AuthTokens `json:"data"`
}

type AuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (ts *TestStruct) AuthTests(t *testing.T) {
	t.Run("SignUp new User", func(t *testing.T) {
		payload, _ := json.Marshal(models.CreateUserInput{
			Email:                "mock@john.com",
			Password:             "123123",
			PasswordConfirmation: "123123",
		})

		req, err := http.NewRequest("POST", "/signup", bytes.NewReader(payload))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ts.Router.ServeHTTP(w, req)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("SignIn Correct User", func(t *testing.T) {
		payload, _ := json.Marshal(models.CreateUserInput{
			Email:    "mock@john.com",
			Password: "123123",
		})

		req, err := http.NewRequest("POST", "/signin", bytes.NewReader(payload))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		ts.Router.ServeHTTP(w, req)

		res := w.Result()

		buf := new(bytes.Buffer)

		buf.ReadFrom(res.Body)

		_ = json.Unmarshal(buf.Bytes(), &ts.Auth)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("SignIn Wrong User", func(t *testing.T) {
		payload, _ := json.Marshal(models.CreateUserInput{
			Email:    "mock@john.com",
			Password: "1231234",
		})

		req, _ := http.NewRequest("POST", "/signin", bytes.NewReader(payload))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ts.Router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}

func (ts *TestStruct) LogoutTest(t *testing.T) {
	t.Run("Logout", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/logout", bytes.NewReader([]byte(nil)))

		ts.SetHeader(req)

		w := httptest.NewRecorder()

		ts.Router.ServeHTTP(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
