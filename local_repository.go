package main

import (
	"os"
	"path/filepath"
	"strings"
)

type LocalRepository struct {
	Path string
}

type LocalRepositoryIndex struct {
	repos []LocalRepository
}

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
		index.repos = append(index.repos, LocalRepository{Path: dir})
	}
	return
}
