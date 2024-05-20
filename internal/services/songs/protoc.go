package songs

import (
	"mime/multipart"
)

type UploadReq struct {
	File *multipart.FileHeader `form:"file" binding:"required,audioOnly"`
}
