package main

import (
	"flag"
	"fmt"
	"go-jvm/src/ch04/classfile"
	"go-jvm/src/ch04/classpath"
	. "go-jvm/src/ch04/utils"
	"os"
	"strings"
)

type cmd struct {
	helpFlag    bool
	versionFlag bool
	classPath   string
	XjreOption  string
	class       string
	args        []string
}

func parseCmd() *cmd {
	cmd := new(cmd)

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "h", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version message")
	flag.StringVar(&cmd.classPath, "classpath", "", "classpath")
	flag.StringVar(&cmd.classPath, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-option] class [args...]\n", os.Args[0])
}

func startJVM(cmd *cmd) {
	classPath := classpath.Parse(cmd.XjreOption, cmd.classPath)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classFile := loadClass(className, classPath)
	fmt.Println(cmd.class)
	printClassInfo(classFile)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	classFile, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return classFile
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constant count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: %v\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.ClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, field := range cf.Fields() {
		fmt.Printf("  %s(%s)\n", field.Name(), field.Description())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, method := range cf.Methods() {
		fmt.Printf("  %s(%s)\n", method.Name(), method.Description())
	}
}

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version %s\n", VERSION)
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
