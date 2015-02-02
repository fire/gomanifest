package main

import (
	"archive/tar"
	"flag"
	"os"
	"path/filepath"
	"log"
	"encoding/json"
	"io/ioutil"
)

var manifest []*tar.Header

func visit(path string, f os.FileInfo, err error) error {
	log.Printf("Visited: %s\n", path)
	h, _ := tar.FileInfoHeader(f, path)
	if err != nil {
        log.Printf("Error: %s", err)
		return err;
	}
	h.Name = filepath.ToSlash(path)
	//log.Println(string(b))
	manifest = append(manifest, h)
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	if root == "" {
		root = "."
	} 
	err := filepath.Walk(root, visit)
	if err != nil {
		log.Printf("filepath.Walk() returned %v\n", err)
	}
	m, err := json.Marshal(manifest)
	if err != nil {
        log.Printf("Error: %s", err)
		panic(err)
	}
	log.Printf("Manifest is: %s", m)
	ioutil.WriteFile("manifest.json", []byte(m), 0x755)
}
