package entry

import (
	"os"
	"path/filepath"
	"strings"
)

func NewWildcardEntry(path string) CompositeEntry {
	baseDir := filepath.Dir(path) // remove *
	entry := CompositeEntry{}
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			subEntry := NewZipEntry(path)
			entry = append(entry, subEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFunc)
	return entry
}
