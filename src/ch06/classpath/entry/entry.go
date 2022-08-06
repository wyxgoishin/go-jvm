package entry

import (
	"os"
	"strings"
)

const PathListSeparator = string(os.PathListSeparator)

type Entry interface {
	ReadClass(className string) ([]byte, Entry, error)
	String() string
}

func NewEntry(path string) Entry {
	if strings.Contains(path, PathListSeparator) {
		return NewCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return NewWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
		return NewZipEntry(path)
	}
	return NewDirEntry(path)
}
