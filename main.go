package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var formats = map[string]string{
	"jpg":  "images",
	"jpeg": "images",
	"png":  "images",
	"txt":  "txt",
	"doc":  "doc",
	"docx": "docx",
	"gif":  "images",
	"py":   "scripts",
	"sql":  "scripts",
	"sh":   "scripts",
	"bash": "scripts",
}

const OUTDIR = "garbage"

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("could not get user's home dir: %v", err)
	}

	fis, err := ioutil.ReadDir(homeDir)
	if err != nil {
		log.Fatalf("could not read files in %s: %v", homeDir, err)
	}

	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		ext := filepath.Ext(fi.Name())
		if ext == "" {
			continue
		}
		ext = ext[1:]

		destination, ok := formats[ext]
		if !ok {
			continue
		}

		destinationDir := filepath.Join(homeDir, OUTDIR, destination)

		err := os.MkdirAll(destinationDir, 0755)
		if err != nil {
			log.Fatalf("could not create destination dir at %s: %v", destinationDir, err)
		}

		err = os.Rename(filepath.Join(homeDir, fi.Name()), filepath.Join(destinationDir, fi.Name()))
		if err != nil {
			log.Fatalf("could not move %s: %v", fi.Name(), err)
		}
	}
}
