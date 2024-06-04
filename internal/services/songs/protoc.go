package songs

import (
	"mime/multipart"
)

type PlayReq struct {
	ID string `uri:"id" `
}

type UploadReq struct {
	File *multipart.FileHeader `form:"file" binding:"required,audioOnly"`
}
