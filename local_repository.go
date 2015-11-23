package main

import (
	"os"
	"path/filepath"
	"strings"
)

type LocalRepository struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (lr LocalRepository) git(args ...string) []string {
	cmdArgs := []string{
		"-C", lr.Path, "--git-dir", "./.git", "--work-tree", "./",
	}
	cmdArgs = append(cmdArgs, args...)
	return cmdArgs
}

func (lr LocalRepository) RunGit(args ...string) error {
	cmd := lr.git(args...)
	return run("git", cmd...)
}

func (lr LocalRepository) CaptureGit(args ...string) (string, error) {
	cmd := lr.git(args...)
	return capture("git", cmd...)
}

type LocalRepositoryIndex []LocalRepository

func NewLocalRepositoryIndexFromPrefix(prefix string) (index LocalRepositoryIndex, err error) {
	pattern := filepath.Join(prefix, "*")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return
	}

	for _, dir := range matches {
		fi, err := os.Stat(dir)
		if err != nil || !fi.IsDir() || strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		index = append(index, LocalRepository{Name: filepath.Base(dir), Path: dir})
	}
	return
}
