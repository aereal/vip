package main

import (
	"os"
	"path/filepath"
	"strings"
)

type LocalRepository struct {
	Path string
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
		index = append(index, LocalRepository{Path: dir})
	}
	return
}
