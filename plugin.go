package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

type PluginName string

type Plugin struct {
	Name PluginName `json:"name"`
	URL  string     `json:"url"`
}

func (plugin *Plugin) Destination() string {
	base := path.Base(plugin.URL)
	return "plugins/" + humanish(base)
}

func NewPlugin(name PluginName, url string) *Plugin {
	return &Plugin{Name: name, URL: url}
}

func humanish(s string) string {
	return strings.Replace(s, ".git", "", 1)
}

// インストールしたいプラグインを持つ
// ユーザが書いた JSON かなにかから作られる予定
type Recipe map[PluginName]*Plugin

func (recipe Recipe) Add(plugin *Plugin) {
	recipe[plugin.Name] = plugin
}

func (recipe Recipe) Size() int {
	return len(recipe)
}

func (recipe Recipe) ByName(name PluginName) *Plugin {
	return recipe[name]
}

func NewRecipe() Recipe {
	return Recipe{}
}

// { "plugins": [{"name": "my-plugin", "url": "https://github.com/you/my-plugin"}, ...] }
type RecipeManifest struct {
	Plugins []*Plugin `json:"plugins"`
}

func NewRecipeFromManifestJSON(path string) Recipe {
	file, e := os.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	dec := json.NewDecoder(file)
	var manifest RecipeManifest
	for {
		if err := dec.Decode(&manifest); err == io.EOF {
			recipe := NewRecipe()
			for _, p := range manifest.Plugins {
				recipe.Add(p)
			}
			return recipe
		} else if err != nil {
			log.Fatal(err)
		}
	}
}
