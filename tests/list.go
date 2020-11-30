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
	ID     uint       `json:"id"`
	Title  string     `json:"title"`
	Status uint       `json:"status"`
	Tasks  []TaskData `json:"tasks"`
}

type Task struct {
	Data TaskData `json:"data"`
}

type TaskData struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status uint   `json:"status"`
}

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
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Get List by ID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/list/"+fmt.Sprint(ts.List.Data.ID), nil)

		ts.SetHeader(req)

		w := httptest.NewRecorder()

		ts.Router.ServeHTTP(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Update List", func(t *testing.T) {
		payload, _ := json.Marshal(models.UpdateListInput{
			Title:  "Shopping Updated",
			Status: 1,
		})

		req, err := http.NewRequest("PUT", "/list/"+fmt.Sprint(ts.List.Data.ID), bytes.NewReader(payload))

		ts.SetHeader(req)

		w := httptest.NewRecorder()

		ts.Router.ServeHTTP(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func (ts *TestStruct) TaskTests(t *testing.T) {
	t.Run("Add Task", func(t *testing.T) {
		payload, _ := json.Marshal(models.CreateTaskInput{
			Title:  "Go to Cinema",
			Status: 1,
			ListId: ts.List.Data.ID,
		})

		req, err := http.NewRequest("POST", "/task", bytes.NewReader(payload))

		ts.SetHeader(req)

		w := httptest.NewRecorder()

		ts.Router.ServeHTTP(w, req)

		res := w.Result()

		buf := new(bytes.Buffer)

		buf.ReadFrom(res.Body)

		_ = json.Unmarshal(buf.Bytes(), &ts.Task)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Update Task", func(t *testing.T) {
		payload, _ := json.Marshal(models.UpdateTaskInput{
			Title:  "Go to Movies",
			Status: 2,
		})

		req, err := http.NewRequest("PUT", "/task/"+fmt.Sprint(ts.Task.Data.ID), bytes.NewReader(payload))

		ts.SetHeader(req)

		w := httptest.NewRecorder()

		ts.Router.ServeHTTP(w, req)

		res := w.Result()

		buf := new(bytes.Buffer)

		buf.ReadFrom(res.Body)

		_ = json.Unmarshal(buf.Bytes(), &ts.Task)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
