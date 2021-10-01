package service

type FileUploadService interface {
	UploadFile() error
	RetrieveFile(location string) error
}
