package model

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	code  int
	Error interface{} `json:"err"`
}

func ErrorResp(err error) *Error {
	return &Error{
		Error: err,
	}
}

func (r *Error) Code() int {
	return http.StatusOK
}

func (r *Error) Encode() ([]byte, error) {
	return json.Marshal(r)
}
