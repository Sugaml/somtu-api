package uploader

import (
	"io"
	"os"
	"path/filepath"
)

type LocalUploader struct {
	BasePath string
}

func NewLocalUploader() *LocalUploader {
	return &LocalUploader{
		BasePath: os.Getenv("FS_LOCATION"),
	}
}

func (u *LocalUploader) UploadFile(file *FileDetails) (string, error) {
	dirPath := filepath.Join(u.BasePath, string(file.FileType), file.EntityID)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return "", err
	}

	destPath := filepath.Join(dirPath, file.FileHeader.Filename)
	dst, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file.File)
	if err != nil {
		return "", err
	}

	return destPath, nil
}

func (u *LocalUploader) GetFileURL(fileType FileType, entityID, fileName string) (string, error) {
	return filepath.Join(u.BasePath, string(fileType), entityID, fileName), nil
}
