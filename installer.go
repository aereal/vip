package main

import (
	"log"
	"os"
)

func Install(plugin *Plugin, c chan int) {
	log.Printf("Install %s ...", plugin.Name)
	_, err := os.Stat(plugin.Destination())
	if err == nil {
		log.Printf("Already exists: %v", plugin.Destination())
		c <- 1
		return
	}
	run("git", "clone", "--depth", "1", plugin.URL, plugin.Destination())
	c <- 1
}

func BatchInstall(recipe Recipe) {
	n := recipe.Size()
	c := make(chan int, n)
	for k := range recipe {
		plugin := recipe.ByName(k)
		go Install(plugin, c)
	}
	for i := 0; i < n; i++ {
		<-c
	}
}
