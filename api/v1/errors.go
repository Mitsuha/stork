package v1

import "errors"

var (
	ServerError  = errors.New("server error")
	Unauthorized = errors.New("unauthorized")
	BadRequest   = errors.New("bad request")
	NotFound     = errors.New("not found")
)
