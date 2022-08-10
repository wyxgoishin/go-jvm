package classpath

import (
	"archive/zip"
	"errors"
	. "go-jvm/src/utils"
	"io/ioutil"
	"log"
	"path/filepath"
)

type ZipEntry struct {
	absDir string
}

func (entry *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	zipReader, err := zip.OpenReader(entry.absDir)
	if err != nil {
		return nil, nil, err
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		if file.Name == className {
			fileReader, err := file.Open()
			if err != nil {
				return nil, nil, err
			}
			defer fileReader.Close()

			data, err := ioutil.ReadAll(fileReader)
			if err != nil {
				return nil, nil, err
			}
			return data, entry, nil
		}
	}

	return nil, nil, errors.New(ErrorClassNotFound)
}

func (entry *ZipEntry) String() string {
	return entry.absDir
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	entry := &ZipEntry{
		absDir: absDir,
	}
	return entry
}
