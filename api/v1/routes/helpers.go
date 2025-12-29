package data_parser

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	// ErrFileNameInvalid represents a file name is invalid.
	ErrFileNameInvalid = "file name is invalid"
	// ErrFileExtensionInvalid represents a file extension is invalid.
	ErrFileExtensionInvalid = "file extension is invalid"
	// ErrFileSizeInvalid represents a file size is invalid.
	ErrFileSizeInvalid = "file size is invalid"
)

var (
	// ErrFileExists represents the file already exists.
	ErrFileExists = errors.New("file already exists")
	// ErrFileNotExist represents the file does not exist.
	ErrFileNotExist = errors.New("file does not exist")
)

// File is a file entity.
type File struct {
	ID        string    `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Path      string    `json:"path" validate:"required"`
	Extension string    `json:"extension" validate:"required"`
	Size      int64     `json:"size" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}

// NewFile creates a new file entity.
func NewFile(id, name, path, extension string, size int64, createdAt, updatedAt time.Time) *File {
	return &File{
		ID:        id,
		Name:      name,
		Path:      path,
		Extension: extension,
		Size:      size,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// Validate validates a file entity.
func (f *File) Validate() error {
	return validate.Struct(f)
}

// GetFileExtension returns the file extension.
func GetFileExtension(file *File) string {
	return filepath.Ext(file.Path)
}

// IsFileExists checks if a file exists.
func IsFileExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, ErrFileNotExist
	}
	return false, err
}

// GetFileSize returns the file size.
func GetFileSize(file string) (int64, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}

// SaveFileToDisk saves a file to disk.
func SaveFileToDisk(file io.Reader, filename string) error {
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}
	return nil
}

// ParseFileSize parses a file size from a string.
func ParseFileSize(sizeStr string) (int64, error) {
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		return 0, err
	}
	if size < 0 {
		return 0, errors.New("file size is invalid")
	}
	return size, nil
}

// IsFileNameValid checks if a file name is valid.
func IsFileNameValid(name string) bool {
	return len(name) > 0 && !strings.HasPrefix(name, ".")
}

// IsFileExtensionValid checks if a file extension is valid.
func IsFileExtensionValid(extension string) bool {
	ext := filepath.Ext(extension)
	return len(ext) > 0 && ext[1:] != ""
}

// NewFileID generates a new random file ID.
func NewFileID() (string, error) {
	return generateRandomString(32)
}

// NewRandomString generates a new random string.
func NewRandomString(length int) string {
	return generateRandomString(length)
}

func generateRandomString(length int) (string, error) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent 64 unique letters
		letterIdxMask = 1 << letterIdxBits - 1
	)

	b := make([]byte, length)
	if length > 0 {
		if rand.Int63()&letterIdxMask == 0 {
			b[0] = letterBytes[rand.Intn(len(letterBytes))]
		}
		for i := 1; i < length; i {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
	}
	return string(b), nil
}

// GetFileBaseName returns the base name of a file.
func GetFileBaseName(file string) string {
	return filepath.Base(file)
}

// GetFileLastModified returns the last modified time of a file.
func GetFileLastModified(file string) (time.Time, error) {
	f, err := os.Open(file)
	if err != nil {
		return time.Time{}, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return time.Time{}, err
	}

	return info.ModTime(), nil
}