package main

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"strings"
)

type PluginName string

type Plugin struct {
	Name PluginName `json:"name"`
	URL  string     `json:"url"`
}

func (plugin *Plugin) Path() string {
	return humanish(path.Base(plugin.URL))
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

func NewRecipeFromManifestJSON(path string) (recipe Recipe, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	dec := json.NewDecoder(file)
	var manifest RecipeManifest
	for {
		if err = dec.Decode(&manifest); err == io.EOF {
			recipe = NewRecipe()
			err = nil
			for _, p := range manifest.Plugins {
				recipe.Add(p)
			}
			return
		} else if err != nil {
			return
		}
	}
}
