package main

import (
	"os"
	"path/filepath"
	"strings"
)

type Receipt struct {
	Directory string
}

type ReceiptIndex struct {
	receipts []Receipt
}

func NewReceiptIndexFromPrefix(prefix string) (index ReceiptIndex, err error) {
	pattern := filepath.Join(prefix, "*")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return
	}

	for _, dir := range matches {
		fi, err := os.Stat(dir)
		if err != nil || !fi.IsDir() || strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		index.receipts = append(index.receipts, Receipt{Directory: dir})
	}
	return
}
