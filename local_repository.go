package main

import (
	"os"
	"path/filepath"
	"strings"
)

type LocalRepository struct {
	Name string
	Path string
}

func (lr LocalRepository) GitDir() string {
	return filepath.Join(lr.Path, ".git")
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
