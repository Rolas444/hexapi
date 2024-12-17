package domain

import (
	"errors"
	"fmt"
)

const (
	ErrCodeDuplicateKey = "duplicate_key"
	ErrCodeInvalidData  = "invalid_data"
	ErrCodeNotFound     = "not_found"
	ErrCodeAccessDenied = "access_denied"
	ErrCodeInsert       = "insert_err"
)

var ErrDuplicateKey = errors.New("duplicate key error")
var ErrInvalidData = errors.New("invalid data")
var ErrNotFound = errors.New("no data found")
var ErrAccessDenied = errors.New("access denied")
var ErrInsert = errors.New("insert error")

type CustomError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Msg)
}
