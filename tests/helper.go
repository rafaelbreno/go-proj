package tests

import (
	"net/http"
)

func (ts *TestStruct) SetHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", ts.Auth.Data.AccessToken)
}
