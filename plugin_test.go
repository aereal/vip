package main

import (
	. "github.com/onsi/gomega"
)
import "testing"

func TestPlugin(t *testing.T) {
	RegisterTestingT(t)

	var (
		plugin *Plugin
	)

	plugin = NewPlugin("test-plugin", "https://github.com/example/test-plugin")
	Expect(plugin.Path()).To(Equal("test-plugin"))

	plugin = NewPlugin("test-plugin", "https://github.com/example/test-plugin.git")
	Expect(plugin.Path()).To(Equal("test-plugin"))
}

func TestNewRecipeFromManifestJSON(t *testing.T) {
	RegisterTestingT(t)

	var (
		recipe *Recipe
		err    error
	)

	recipe, err = NewRecipeFromManifestJSON("test/assets/valid-plugins.json")
	Expect(err).To(BeNil())
	Expect(recipe).NotTo(BeNil())

	recipe, err = NewRecipeFromManifestJSON("test/assets/broken.json")
	Expect(err).NotTo(BeNil())
	Expect(recipe).To(BeNil())
}
