package main

import (
	"fmt"
	. "go-jvm/src/utils"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var workPath = "/Users/yixinwu/go-jvm/src"
var src = "frame.operandStack"
var dst = "frame.operandStack()"

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
				if !info.IsDir() && !strings.Contains(info.Name(), "refractor.go") {
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
