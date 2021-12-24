package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	"pdf":  "pdf",
}

const OUTDIR = "garbage"

func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("could not get user's home dir: %v", err)
	}

	targetDirPtr := flag.String("targetDir", homeDir, "determine where the janitor should clean the files")
	flag.Parse()

	targetDir := *targetDirPtr
	if !filepath.IsAbs(targetDir) {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("could not get cwd: %v", err)
		}
		targetDir = filepath.Join(wd, targetDir)
	}

	fis, err := ioutil.ReadDir(targetDir)
	if err != nil {
		log.Fatalf("could not read files in %s: %v", *targetDirPtr, err)
	}

	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(fi.Name()))
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

		err = os.Rename(filepath.Join(targetDir, fi.Name()), filepath.Join(destinationDir, fi.Name()))
		if err != nil {
			log.Fatalf("could not move %s: %v", fi.Name(), err)
		}
	}
}
