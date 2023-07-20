package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

// ExtractPathFromAbsolutePath returns path without file name from absolute path
func ExtractPathFromAbsolutePath(absolutePath string) string {
	if absolutePath == "" {
		return absolutePath
	}
	pos := strings.LastIndex(absolutePath, "/")
	if pos >= 0 {
		return absolutePath[:pos]
	}
	return absolutePath
}

// ExtractFileNameFromAbsolutePath returns file name from absolute path
func ExtractFileNameFromAbsolutePath(absolutePath string) string {
	if absolutePath == "" {
		return absolutePath
	}
	pos := strings.LastIndex(absolutePath, "/")
	if pos >= 0 {
		return absolutePath[pos+1:]
	}
	return absolutePath
}

// ExtractFileExtensionFromAbsolutePath returns file extension from absolute path
func ExtractFileExtensionFromAbsolutePath(absolutePath string) string {
	if absolutePath == "" {
		return absolutePath
	}
	pos := strings.LastIndex(absolutePath, ".")
	if pos >= 0 {
		return absolutePath[pos+1:]
	}
	return absolutePath
}

// ReadFileContent returns file content as bytes slice
func ReadFileContent(absolutePath string) ([]byte, error) {
	var read []byte
	if absolutePath == "" {
		return read, nil
	}
	f, err := os.Open(absolutePath)

	if err != nil {
		return read, err
	}
	defer f.Close()

	read, err = ioutil.ReadAll(f)

	if err != nil {
		return read, err
	}

	return read, nil
}

// CreateDirectory creates directory with given path
func CreateDirectory(absolutePath string) error {
	if absolutePath == "" {
		return nil
	}
	_, err := os.Stat(absolutePath)
	if err != nil {
		err = os.Mkdir(absolutePath, 0777)
	}

	return err
}

// DeleteDirectory deletes directory with given path
func DeleteDirectory(absolutePath string) error {
	if absolutePath == "" {
		return nil
	}
	_, err := os.Stat(absolutePath)
	if err != nil {
		return err
	}
	err = os.RemoveAll(absolutePath)
	if err != nil {
		return err
	}

	return err
}

// WriteContentToFile returns file with content written, creates file if it not exist
func WriteContentToFile(absolutePath string, bytes []byte) error {
	if absolutePath == "" {
		return errors.New("Empty path given")
	}

	f, err := os.OpenFile(absolutePath,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0777)

	if err != nil {
		return err
	}

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

// TrimSuffix removes given suffix from given string
func TrimSuffix(absolutePath string, suffix string) string {
	if strings.HasSuffix(absolutePath, suffix) {
		return strings.TrimSuffix(absolutePath, suffix)
	}
	return absolutePath
}

// SplitPathToSegments splits given path in format /path/to/file into string slice [path, to, file]
func SplitPathToSegments(absolutePath string, separator string) []string {
	if absolutePath == "" {
		return []string{}
	}

	return strings.Split(absolutePath, separator)

}

// CheckIfPathExist returns true if path exist, false otherwise
func CheckIfPathExist(absolutePath string) bool {
	_, err := os.Stat(absolutePath)
	if err != nil {
		return false
	}
	return true
}

// CheckIfIsDir returns true if path is directory otherwise false. To check if path exist use CheckIfPathExist
func CheckIfIsDir(absolutePath string) bool {
	fi, _ := os.Stat(absolutePath)
	if fi != nil {
		return fi.IsDir()
	}
	return false
}

// RenameFile renames given path with new file name
func RenameFile(absolutePath string, newName string) error {
	if absolutePath == "" {
		return errors.New("Empty path given")
	}

	return os.Rename(absolutePath, newName)
}
