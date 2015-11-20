package main

import (
	"fmt"
	"sync"
)

type Environment []*Deployment

func NewEnvironment(prefixPath string) (Environment, error) {
	var env Environment
	var err error

	index, err := NewLocalRepositoryIndexFromPrefix(prefixPath)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	for _, localrepo := range index {
		wg.Add(1)
		go func(repo LocalRepository) {
			defer wg.Done()
			deploy, err := NewDeploymentFromLocalRepository(repo)
			if err != nil {
				return
			}
			env = append(env, deploy)
		}(localrepo)
	}
	wg.Wait()

	return env, nil
}

type Deployment struct {
	repo     *LocalRepository
	revision string
}

func NewDeploymentFromLocalRepository(repo LocalRepository) (*Deployment, error) {
	rev, err := capture("git", "--git-dir", repo.GitDir(), "rev-parse", "HEAD")
	if err != nil {
		return nil, err
	}
	return &Deployment{repo: &repo, revision: rev}, nil
}

func (d *Deployment) Format() string {
	return fmt.Sprintf("%s (revision: %s)", d.repo.Name, d.revision)
}
