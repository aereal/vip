package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func (e Environment) Lock(lockFileName string) (err error) {
	file, err := os.OpenFile(lockFileName, os.O_WRONLY|os.O_CREATE, 0666)
	enc := json.NewEncoder(file)
	err = enc.Encode(e)
	return
}

type Deployment struct {
	Repo     *LocalRepository `json:"repo"`
	Revision string           `json:"rev"`
}

func NewDeploymentFromLocalRepository(repo LocalRepository) (*Deployment, error) {
	rev, err := capture("git", "--git-dir", repo.GitDir(), "rev-parse", "HEAD")
	if err != nil {
		return nil, err
	}
	return &Deployment{Repo: &repo, Revision: rev}, nil
}

func (d *Deployment) Format() string {
	return fmt.Sprintf("%s (revision: %s)", d.Repo.Name, d.Revision)
}
