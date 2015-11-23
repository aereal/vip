package main

import (
	. "github.com/onsi/gomega"
)
import "testing"

func TestNewLocalRepositoryIndexFromPrefix(t *testing.T) {
	RegisterTestingT(t)

	var (
		index LocalRepositoryIndex
		err   error
	)

	index, err = NewLocalRepositoryIndexFromPrefix("test/assets/bundle/")
	Expect(index).NotTo(BeNil())
	Expect(index[0].Name).To(Equal("vim-plugin-1"))
	Expect(index[0].Path).To(Equal("test/assets/bundle/vim-plugin-1"))
	Expect(err).To(BeNil())
}
