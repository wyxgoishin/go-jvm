package classpath

import (
	. "go-jvm/src/ch02/classpath/entry"
	. "go-jvm/src/ch02/utils"
	"os"
	"path/filepath"
	"strings"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, classpath string) *Classpath {
	cp := new(Classpath)
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(classpath)
	return cp
}

func (cp *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className += ".class"
	if data, entry, err := cp.bootClasspath.ReadClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := cp.extClasspath.ReadClass(className); err == nil {
		return data, entry, nil
	}
	return cp.userClasspath.ReadClass(className)
}

func (cp *Classpath) String() string {
	var strs []string
	strs = append(strs, cp.bootClasspath.String())
	strs = append(strs, cp.extClasspath.String())
	strs = append(strs, cp.userClasspath.String())
	return strings.Replace(strings.Join(strs, PathListSeparator), strings.Repeat(PathListSeparator, 2), PathListSeparator, -1)
}

func (cp *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClasspath = NewWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	cp.extClasspath = NewWildcardEntry(jreExtPath)
}

func (cp *Classpath) parseUserClasspath(classpath string) {
	// 如果没有提供cp，则默认以当前目录作为cp
	if classpath == "" {
		classpath = "."
	}
	cp.userClasspath = NewEntry(classpath)
}

func getJreDir(jreOption string) string {
	if Exists(jreOption) {
		return jreOption
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("can not find jre folder")
}
