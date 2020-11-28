package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-proj/app/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type List struct {
	Data ListData `json:"data"`
}

type ListData struct {
	ID    uint    `json:"id"`
	Title string  `json:"title"`
	Tasks []Tasks `json:"tasks"`
}

type Tasks struct{}

func (ts *TestStruct) ListTests(t *testing.T) {

	// Execute a test
	// t.Run("Test Title", func () { performs test })
	t.Run("Success List", func(t *testing.T) {
		// Create a """"""""json body""""""""
		payload, _ := json.Marshal(models.CreateListInput{
			Title:  "Shopping List",
			Status: 1,
		})

		// Create a new Request
		req, err := http.NewRequest("POST", "/list", bytes.NewReader(payload))

		// Setting Default Headers
		ts.SetHeader(req)

		w := httptest.NewRecorder()

		ts.Router.ServeHTTP(w, req)

		res := w.Result()

		buf := new(bytes.Buffer)

		buf.ReadFrom(res.Body)

		_ = json.Unmarshal(buf.Bytes(), &ts.List)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Error List", func(t *testing.T) {
		// Create a new Request
		req, err := http.NewRequest("POST", "/list", nil)

		// Setting Default Headers
		ts.SetHeader(req)

		w := httptest.NewRecorder()

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Get List by ID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/list/"+fmt.Sprint(ts.List.Data.ID), nil)

		ts.SetHeader(req)

		w := httptest.NewRecorder()

		ts.Router.ServeHTTP(w, req)

		t.Log(err)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
