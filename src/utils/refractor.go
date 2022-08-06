package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var workPath = "/Users/xxx/go-jvm/src/ch11"
var src = "ch10"
var dst = "ch11"

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsFile(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && !stat.IsDir()
}

func main() {
	var filesToRefractor []string
	if Exists(workPath) {
		if IsFile(workPath) {
			filesToRefractor = append(filesToRefractor, workPath)
		} else {
			filepath.Walk(workPath, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					filesToRefractor = append(filesToRefractor, path)
				}
				return nil
			})
		}
		for _, fp := range filesToRefractor {
			refractor(fp, src, dst)
			fmt.Printf("Refractor %q to %q in %q\n", src, dst, fp)
		}
	}

}

func refractor(filepath, src, dst string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic("cannot read " + filepath)
	}

	content := strings.Replace(string(data), src, dst, -1)
	ioutil.WriteFile(filepath, []byte(content), 0755)
}
