package v1

import "errors"

var (
	ServerError  = errors.New("server error")
	Unauthorized = errors.New("unauthorized")
)
