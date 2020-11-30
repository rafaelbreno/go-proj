package main

import (
	"go-proj/routes"
	"go-proj/tests"
	"testing"
)

var MainTest tests.TestStruct

func TestAuth(t *testing.T) {
	MainTest.Router = routes.GetTestRouter()

	MainTest.AuthTests(t)

	MainTest.ListTests(t)

	MainTest.TaskTests(t)

	MainTest.LogoutTest(t)

}
