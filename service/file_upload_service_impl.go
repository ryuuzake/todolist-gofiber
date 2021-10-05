package service

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"mime/multipart"
)

type FileUploadServiceImpl struct {
}

func (service FileUploadServiceImpl) UploadFile(file *multipart.FileHeader) (string, error) {
	err := fasthttp.SaveMultipartFile(file, fmt.Sprintf("./storage/%s", file.Filename))

	return fmt.Sprintf("http://127.0.0.1:3000/storage/%s", file.Filename), err
}

func (service FileUploadServiceImpl) RetrieveFile(location string) (string, error) {
	return fmt.Sprintf("./storage/%s", location), nil
}
