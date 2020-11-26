package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-proj/app/models"
	"go-proj/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Test struct{}

func (_ *Test) AuthTests(t *testing.T) {
	r := routes.GetTestRouter()

	t.Run("SignUp new User", func(t *testing.T) {
		payload, _ := json.Marshal(models.CreateUserInput{
			Email:                "mock@john.com",
			Password:             "123123",
			PasswordConfirmation: "123123",
		})

		req, err := http.NewRequest("POST", "/signup", bytes.NewReader(payload))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
