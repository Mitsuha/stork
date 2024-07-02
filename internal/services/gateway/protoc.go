package gateway

type FileRequest struct {
	Filename string `uri:"filename" binding:"required"`
}
