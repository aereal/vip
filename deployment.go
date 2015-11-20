package main

import (
	"fmt"
)

type Environment []*Deployment

func NewEnvironment(prefixPath string) (Environment, error) {
	var env Environment
	var err error

	index, err := NewLocalRepositoryIndexFromPrefix(prefixPath)
	if err != nil {
		return nil, err
	}
	for _, localrepo := range index {
		deploy, err := NewDeploymentFromLocalRepository(localrepo)
		if err != nil {
			continue
		}
		env = append(env, deploy)
	}

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
