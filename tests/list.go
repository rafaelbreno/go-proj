package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-proj/app/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type List struct{}

type Item struct{}

func (ts *TestStruct) ListTests(t *testing.T) {
	// Execute a test
	// t.Run("Test Title", func () { performs test })
	t.Run("Success List", func(t *testing.T) {
		// Create a """"""""json body""""""""
		payload, _ := json.Marshal(models.CreateListInput{
			Title:  "Shopping List",
			Status: 0,
		})

		// Create a new Request
		req, err := http.NewRequest("POST", "/task", bytes.NewReader(payload))

		// Setting Default Headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", ts.Data.AccessToken)

		w := httptest.NewRecorder()

		ts.Router.ServeHTTP(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
