package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TODO: move this to yaml so users can configure it and I dont have to recompile every time
var formats = map[string]string{
	"jpg":     "images",
	"jpeg":    "images",
	"png":     "images",
	"txt":     "txt",
	"doc":     "doc",
	"docx":    "doc",
	"gif":     "images",
	"py":      "scripts",
	"sql":     "scripts",
	"sh":      "scripts",
	"bash":    "scripts",
	"pdf":     "pdf",
	"csv":     "csv",
	"zip":     "trash",
	"xls":     "doc",
	"xlsx":    "doc",
	"xlsm":    "doc",
	"xml":     "other",
	"html":    "other",
	"pkg":     "trash",
	"yaml":    "scripts",
	"yml":     "scripts",
	"conf":    "scripts",
	"msg":     "doc",
	"mov":     "video",
	"json":    "scripts",
	"parquet": "doc",
	"gz":      "trash",
	"webp":    "video",
	"pptx":    "doc",
	"ppt":     "doc",
	"svg":     "images",
	"mp4":     "video",
	"app":     "trash",
	"dmg":     "trash",
	"log":     "log",
	"out":     "log",
	"diff":    "trash",
	"ipynb":   "scripts",
	"heic":    "images",
	"numbers": "trash",
	"webm":    "video",
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

	fmt.Printf("cleaning directory: `%s`\n", targetDir)

	fis, err := ioutil.ReadDir(targetDir)
	if err != nil {
		log.Fatalf("could not read files in %s: %v", *targetDirPtr, err)
	}

	for _, fi := range fis {
		ext := strings.ToLower(filepath.Ext(fi.Name()))
		if ext == "" {
			continue
		}

		if fi.IsDir() && ext != ".app" {
			continue
		}

		ext = ext[1:]

		destination, ok := formats[ext]
		if !ok {
			fmt.Printf("no destination for: `%s`\n", ext)
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
