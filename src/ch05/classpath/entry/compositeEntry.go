package entry

import (
	"fmt"
	"os"
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
	return nil, nil, fmt.Errorf("no such class '%v' in '%v'", className, entry.String())
}

func (entry CompositeEntry) String() string {
	strs := make([]string, len(entry))
	for idx, subEntry := range entry {
		strs[idx] = subEntry.String()
	}
	return strings.Join(strs, PathListSeparator)
}

func NewCompositeEntry(pathList string) *CompositeEntry {
	entry := CompositeEntry{}
	for _, path := range strings.Split(pathList, string(os.PathListSeparator)) {
		subEntry := NewEntry(path)
		entry = append(entry, subEntry)
	}
	return &entry
}
