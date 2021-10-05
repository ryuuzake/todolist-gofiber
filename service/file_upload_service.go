package service

import (
	"mime/multipart"
)

type FileUploadService interface {
	UploadFile(file *multipart.FileHeader) (string, error)
	RetrieveFile(location string) (string, error)
}
