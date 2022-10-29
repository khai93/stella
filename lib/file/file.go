package filelib

import (
	"archive/tar"
	"bytes"
)

// Create a buffer representing a file from string
func CreateFileBuffer(content string, fileName string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	err := tw.WriteHeader(&tar.Header{
		Name: fileName,
		Mode: 0777,
		Size: int64(len(content)),
	})
	if err != nil {
		return nil, err
	}

	// write to buffer
	tw.Write([]byte(content))
	tw.Close()

	return &buf, nil
}
