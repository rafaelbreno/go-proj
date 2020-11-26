package main

import (
	"go-proj/tests"
	"testing"
)

var MainTest tests.Test

func TestAuth(t *testing.T) {
	MainTest.AuthTests(t)
}
