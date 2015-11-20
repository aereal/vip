package main

import (
	"fmt"
)

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
