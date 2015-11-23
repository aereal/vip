package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

func install(plugin *Plugin, pathPrefix string) {
	log.Printf("Install %s ...", plugin.Name)
	dest := filepath.Join(pathPrefix, plugin.Path())
	_, err := os.Stat(dest)
	if err == nil {
		log.Printf("Already exists: %v", dest)
		return
	}
	run("git", "clone", "--depth", "1", plugin.URL, dest)
}

func BatchInstall(recipe *Recipe, pathPrefix string) {
	plugins := recipe.Plugins()
	var wg sync.WaitGroup
	for _, plugin := range plugins {
		wg.Add(1)
		go func(p *Plugin) {
			defer wg.Done()
			install(p, pathPrefix)
		}(plugin)
	}
	wg.Wait()
}

func Checkout(env Environment) error {
	deployments := env
	var wg sync.WaitGroup
	for _, deploy := range deployments {
		wg.Add(1)
		go func(d *Deployment) {
			defer wg.Done()
			d.Repo.RunGit("checkout", "-q", d.Revision)
		}(deploy)
	}
	wg.Wait()
	return nil
}
