package utils

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

// GitRepoRemote gets the remote URL of the git repo
func GitRepoRemote(location string) (string, error) {
	path := filepath.Join(location, git.GitDirName)

	r, err := git.PlainOpenWithOptions(path, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return "", err
	}

	rem, err := r.Remote(git.DefaultRemoteName)
	if err != nil {
		return "", err
	}

	return rem.Config().URLs[0], nil
}
