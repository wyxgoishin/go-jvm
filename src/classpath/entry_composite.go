package classpath

import (
	"errors"
	. "go-jvm/src/utils"
	"strings"
)

type CompositeEntry []Entry

func (entry CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, subEntry := range entry {
		data, from, err := subEntry.ReadClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New(ErrorClassNotFound)
}

func (entry CompositeEntry) String() string {
	strs := make([]string, len(entry))
	for idx, subEntry := range entry {
		strs[idx] = subEntry.String()
	}
	return strings.Join(strs, PathListSeparator)
}

func newCompositeEntry(pathList string) *CompositeEntry {
	entry := CompositeEntry{}
	for _, path := range strings.Split(pathList, PathListSeparator) {
		subEntry := newEntry(path)
		entry = append(entry, subEntry)
	}
	return &entry
}
