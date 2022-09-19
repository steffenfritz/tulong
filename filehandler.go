package tulong

import "os"

// Filehandle return a file handle. It has to be closed manually.
func Filehandle(filePath string) (*os.File, error) {
	fd, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	return fd, nil

}
