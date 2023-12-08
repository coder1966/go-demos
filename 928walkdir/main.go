package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {

	walkDir

}

func walkDir(path string, chans chan dataChan, regslice []string) {
	entries, _ := dirents(path)
	for _, e := range entries {
		if e.IsDir() {

			walkDir(filepath.Join(path, e.Name()), chans, regslice)
		} else {
			flag := isreg(filepath.Join(path, e.Name()), regslice)
			if !flag {
				chans <- dataChan{
					fileSize: e.Size(),
				}
			}
		}
	}
}

func dirents(path string) ([]os.FileInfo, bool) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	return entries, true
}
