package main

import (
	"go-proj/tests"
	"testing"
)

var MainTest tests.TestStruct

func TestAuth(t *testing.T) {
	MainTest.AuthTests(t)
}
