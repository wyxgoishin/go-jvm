package entry

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewEntry(t *testing.T) {
	Convey("DirEntry", t, func() {
		Convey("AbsolutePath", func() {
			path := "/Users/xxx/go-jvm/src/ch02/classpath/resources"
			expected, _ := filepath.Abs(path)

			entry := NewEntry(path)
			So(entry.String(), ShouldEqual, expected)

			className := "test.class"
			data, _, err := entry.ReadClass(className)
			So(err, ShouldBeNil)
			realData, _ := ioutil.ReadFile(filepath.Join(expected, className))
			So(data, ShouldResemble, realData)
		})
		Convey("WithoutPath", func() {
			path := ""
			expected := "/Users/xxx/go-jvm/src/ch02/classpath/entry"

			entry := NewEntry(path)
			So(entry.String(), ShouldEqual, expected)

			className := "test.class"
			data, _, err := entry.ReadClass(className)
			So(err, ShouldNotBeNil)
			So(data, ShouldBeNil)
		})
	})
	Convey("CompositeEntry", t, func() {
		path := "/Users/xxx/go-jvm/src/ch02/classpath/resources:/Users/xxx/go-jvm/src/ch02/classpath/resources/cp"
		pathes := strings.Split(path, PathListSeparator)
		for idx, subPath := range pathes {
			absPath, _ := filepath.Abs(subPath)
			pathes[idx] = absPath
		}
		expected := strings.Join(pathes, PathListSeparator)

		entry := NewEntry(path)
		So(entry.String(), ShouldEqual, expected)

		className := "test.class"
		data, _, err := entry.ReadClass(className)
		So(err, ShouldBeNil)
		realData, _ := ioutil.ReadFile(filepath.Join(pathes[0], className))
		So(data, ShouldResemble, realData)
	})
}
