package entry

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func (entry *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(entry.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, entry, err
}

func (entry *DirEntry) String() string {
	return entry.absDir
}

func NewDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	entry := &DirEntry{
		absDir: absDir,
	}
	return entry
}
