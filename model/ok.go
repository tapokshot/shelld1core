package model

import (
	"encoding/json"
	"net/http"
)

type OkResponse struct {
	Data interface{} `json:"data"`
}

func Ok(data interface{}) *OkResponse {
	return &OkResponse{
		Data: data,
	}
}

func (r *OkResponse) Code() int {
	return http.StatusOK
}

func (r *OkResponse) Encode() ([]byte, error) {
	return json.Marshal(r)
}