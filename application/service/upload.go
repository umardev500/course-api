package service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type UploadService struct{}

// NewUploadService create upload service instance
func NewUploadService() *UploadService {
	return &UploadService{}
}

// UploadChunk - is create upload as chunk
func (us *UploadService) UploadChunk(formFile *multipart.FileHeader, dirID string, chunkTotal, chunkIndex int) {
	file, err := formFile.Open()
	if err != nil {
		return
	}
	defer file.Close()

	// check and create folder if not exist
	tempDir := filepath.Join("./temp", dirID)
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		os.MkdirAll(tempDir, os.ModePerm)
	}

	us.Upload(file, tempDir, formFile.Filename)
}

// Upload is service to upload file
//
// Params:
//   - file multipart.File
//   - directory string
//   - filename string
//
// Return:
//   - error
func (us *UploadService) Upload(file multipart.File, directory, filename string) error {
	if _, err := os.Stat(directory); err != nil {
		os.MkdirAll(directory, os.ModePerm)
	}

	fileLocation := filepath.Join(directory, filename)
	out, err := os.Create(fileLocation)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}
