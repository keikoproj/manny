package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	ErrDestinationIsDir  = "Destination is a directory: "
	ErrDestinationExists = "Destination already exists: "
)

// ValidateAndWrite checks that a location does not exist and that it is not a directory
func ValidateAndWrite(location string, bytes []byte) (bool, error) {
	// validation
	location = filepath.Clean(location)
	if fl, err := os.Stat(location); !os.IsNotExist(err) {
		if fl.Mode().IsDir() {
			return false, errors.New(ErrDestinationIsDir + location)
		}

		return false, errors.New(ErrDestinationExists + location)
	}

	// write
	err := ioutil.WriteFile(location, bytes, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}
