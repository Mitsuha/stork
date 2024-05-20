package validator

import (
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	"strings"
)

func AudioOnly(fl validator.FieldLevel) bool {
	file, ok := fl.Field().Interface().(multipart.FileHeader)

	if !ok {
		return false
	}

	return strings.HasPrefix(file.Header.Get("Content-Type"), "audio/")
}
