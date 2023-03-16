package writer

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	// Delimiter default ansdoc delimiter.
	Delimiter = "<!--ansdoc -->"
)

// SplitFile will return the parts of a given file splitted by start and end delimiter.
func SplitFile(filePath string, delimiter string) (string, string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", "", err
	}

	defer f.Close()

	out, err := io.ReadAll(f)
	if err != nil {
		return "", "", err
	}

	parts := strings.Split(string(out), delimiter)

	if len(parts) != 3 {
		return "", "", fmt.Errorf("cannot split %s in 2 parts by delimiter: \"%s\"", filePath, delimiter)
	}

	return parts[0], parts[len(parts)-1], nil
}

// CopyFile copies a src file to ist dest location.
func CopyFile(src, dest string) error {
	_, err := os.Stat(src)
	if err != nil {
		return err
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}

	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destination.Close()

	if _, err := io.Copy(destination, source); err != nil {
		return err
	}

	return nil
}
