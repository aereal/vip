package main

import (
	"log"
	"os"
)

func install(plugin *Plugin, c chan int) {
	log.Printf("Install %s ...", plugin.Name)
	dest := "plugins/" + plugin.Path()
	_, err := os.Stat(dest)
	if err == nil {
		log.Printf("Already exists: %v", dest)
		c <- 1
		return
	}
	run("git", "clone", "--depth", "1", plugin.URL, dest)
	c <- 1
}

func BatchInstall(recipe Recipe) {
	n := recipe.Size()
	c := make(chan int, n)
	for k := range recipe {
		plugin := recipe.ByName(k)
		go install(plugin, c)
	}
	for i := 0; i < n; i++ {
		<-c
	}
}
