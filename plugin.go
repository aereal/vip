package main

import (
	"path"
)

type PluginName string

type Plugin struct {
	Name PluginName
	URL  string
}

func (plugin *Plugin) Destination() string {
	base := path.Base(plugin.URL)
	return "plugins/" + base
}

func NewPlugin(name PluginName, url string) *Plugin {
	return &Plugin{Name: name, URL: url}
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
