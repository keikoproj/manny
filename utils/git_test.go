package utils

import (
	"strings"
	"testing"
)

func Test_GitRepoRemote(t *testing.T) {
	// test table
	tt := []struct {
		Identifier string
		Path       string
		Output     string
		WantErr    bool
		ErrPrefix  string
	}{
		{
			"Valid",
			"../",
			"manny.git",
			false,
			"",
		},
		{
			"Not a git directory",
			"/tmp",
			"",
			true,
			"repository does not exist",
		},
	}

	// testing loop
	for _, tc := range tt {
		t.Run(tc.Identifier, func(t *testing.T) {
			out, err := GitRepoRemote(tc.Path)
			haveErr := err != nil

			// evaluate output
			if !strings.Contains(out, tc.Output) {
				t.Errorf("Error with output: got: %s, want: %s", out, tc.Output)
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
		})
	}
}
