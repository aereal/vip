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
