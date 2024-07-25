package utils

import (
	"encoding/json"
	"net/http"
)

// 响应状态码, 以JSON格式返回数据v
func WriteJSON[T any](w http.ResponseWriter, status int, v T) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}

// 从请求中获取payload
func DecodeFromRequst[T any](r *http.Request) (T, error) {
	var v T

	err := json.NewDecoder(r.Body).Decode(&v)

	return v, err
}

type Message struct {
	Message string `json:"message"`
}
