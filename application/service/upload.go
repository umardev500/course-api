package service

import (
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UploadService struct{}

// NewUploadService create upload service instance
func NewUploadService() *UploadService {
	return &UploadService{}
}

// UploadChunk - is create upload as chunk
func (us *UploadService) UploadChunk(formFile *multipart.FileHeader, dirID string, chunkTotal, chunkIndex int) (finish bool, fileLocation *string, err error) {
	file, err := formFile.Open()
	if err != nil {
		return false, nil, err
	}
	defer file.Close()

	// check and create folder if not exist
	tempDir := filepath.Join("./temp", dirID)
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		os.MkdirAll(tempDir, os.ModePerm)
	}
	err = us.Upload(file, tempDir, formFile.Filename)
	if err != nil {
		return false, nil, err
	}

	if chunkTotal == chunkIndex {
		destDir := filepath.Join("public", "upload")
		fileLocation, err := us.reassembleChunk(tempDir, destDir)
		if err != nil {
			return true, nil, err
		}
		return true, fileLocation, nil
	}

	return false, nil, nil
}

// reassembleChunk reassemble the chunk into singgle file
//
// Params:
//   - tempDir string - temporary dir of chunks
//   - dest string - file path destination for assemble
func (us *UploadService) reassembleChunk(tempDir, dest string) (*string, error) {
	var output []byte
	var fileFormat string

	filepath.Walk(tempDir, func(path string, info fs.FileInfo, failed error) error {
		if failed != nil {
			return failed
		}
		if path == tempDir {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		fileByte, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		output = append(output, fileByte...)
		if fileFormat == "" {
			format := strings.Split(info.Name(), ".")
			wordLen := len(format)
			fileFormat = format[wordLen-1]
		}

		// remove file
		os.Remove(path)

		return nil
	})
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		os.MkdirAll(dest, os.ModePerm)
	}
	now := time.Now().Unix()
	filename := fmt.Sprintf("%s/%d.%s", dest, now, fileFormat)
	out, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	_, err = out.Write(output)
	if err != nil {
		return nil, err
	}

	return &filename, nil
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
