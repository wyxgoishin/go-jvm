package classpath

import (
	. "go-jvm/src/utils"
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := filepath.Dir(path) // remove *
	entry := CompositeEntry{}
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, JarExtension) || strings.HasSuffix(path, strings.ToUpper(JarExtension)) {
			subEntry := newZipEntry(path)
			entry = append(entry, subEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFunc)
	return entry
}
