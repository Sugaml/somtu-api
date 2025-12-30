package uploader

import (
	"fmt"
	"mime/multipart"
	"os"
	"time"
)

// FileType represents the category of the uploaded file
type FileType string

const (
	FileTypeStudentPhoto FileType = "student_photo"
	FileTypeBookPhoto    FileType = "book_photo"
	FileTypeIDCard       FileType = "id_card"
	FileTypeDocument     FileType = "document"
)

type FileMetadata struct {
	FileName    string            // Clean name of the file
	Size        int64             // File size in bytes
	ContentType string            // MIME type
	UploadedAt  time.Time         // Timestamp of upload
	Tags        map[string]string // Optional custom tags
}

// FileDetails represents an uploadable file with associated metadata
type FileDetails struct {
	FileType   FileType
	EntityID   string // studentID, bookID, etc.
	File       multipart.File
	FileHeader *multipart.FileHeader
	Metadata   FileMetadata
}

type FileUploader interface {
	UploadFile(file *FileDetails) (string, error)
	GetFileURL(fileType FileType, entityID, fileName string) (string, error)
}

func GetUploader() (FileUploader, error) {
	fsType := os.Getenv("FS_TYPE")
	switch fsType {
	case "s3":
		return NewS3Uploader()
	case "local":
		return NewLocalUploader(), nil
	default:
		return nil, fmt.Errorf("unsupported FS_TYPE: %s", fsType)
	}
}
