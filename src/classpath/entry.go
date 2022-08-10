package classpath

import (
	. "go-jvm/src/utils"
	"strings"
)

type Entry interface {
	ReadClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, PathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, JarExtension) || strings.HasSuffix(path, strings.ToUpper(JarExtension)) {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
