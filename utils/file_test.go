package utils

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_ValidateAndWrite(t *testing.T) {
	// test table
	tt := []struct {
		// Name of the test
		Identifier string
		// PreDir will create a directory before ValidateAndWrite() is called
		PreDir string
		// PreWrite will write to a file before ValidateAndWrite() is called
		PreWrite string
		// File location to write to, passed to ValidateAndWrite()
		Location string
		// File content to write, passed to ValidateAndWrite()
		Content []byte
		// Whether we want an error
		WantErr bool
		// What you want the return bool to be
		Output bool
		// The error that the message should be prefixed with
		ErrPrefix string
	}{
		{
			"Valid",
			"",
			"",
			"/tmp/manny_test",
			[]byte("test"),
			false,
			true,
			"",
		},
		{
			"Location already exists",
			"",
			"/tmp/manny_test",
			"/tmp/manny_test",
			[]byte("test"),
			true,
			false,
			ErrDestinationExists,
		},
		{
			"Location is a directory",
			"/tmp/manny_test",
			"",
			"/tmp/manny_test",
			[]byte("test"),
			true,
			false,
			ErrDestinationIsDir,
		},
	}

	// testing loop
	for _, tc := range tt {
		var cleanup []string

		t.Run(tc.Identifier, func(t *testing.T) {
			// creates a file
			if tc.PreWrite != "" {
				cleanup = append(cleanup, tc.PreWrite)
				err := ioutil.WriteFile(tc.PreWrite, tc.Content, 0644)
				if err != nil {
					t.Errorf("Error reading file: %s", err)
				}
			}

			// creates a directory
			if tc.PreDir != "" {
				cleanup = append(cleanup, tc.PreDir)
				err := os.Mkdir(tc.PreDir, 0644)
				if err != nil {
					t.Errorf("Error reading file: %s", err)
				}
			}

			got, err := ValidateAndWrite(tc.Location, tc.Content)
			haveErr := err != nil

			// clean up any files that were generated
			if got {
				cleanup = append(cleanup, tc.Location)
			}

			// evaluate output
			if tc.Output != got {
				t.Errorf("Error with output: got: %t, want: %t", got, tc.Output)
			}

			// err prefix is wrong
			if haveErr && tc.WantErr && !strings.HasPrefix(err.Error(), tc.ErrPrefix) {
				t.Errorf("Error not prefixed as expected. got: %s, want: %s", err.Error(), tc.ErrPrefix)
			}

			// didn't expect an error but got one
			if haveErr && !tc.WantErr {
				t.Errorf("Got: %s, want: nil", err)
			}

			// expected an error but didn't get one
			if haveErr && err == nil {
				t.Errorf("Got: nil,")
			}

			b, err := ioutil.ReadFile(tc.Location)
			if err != nil && tc.Output {
				t.Errorf("Error reading file: %s", err)
			}

			if !reflect.DeepEqual(tc.Content, b) && tc.Output {
				t.Errorf("File content is not the same. got: %s, want: %s", b, tc.Content)
			}
		})

		for _, location := range cleanup {
			err := os.Remove(location)
			if err != nil {
				t.Errorf("Error removing file: %s", err)
			}
		}
	}
}
