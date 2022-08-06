package main

import (
	"flag"
	"fmt"
	"go-jvm/src/ch06/classpath"
	"go-jvm/src/ch06/rtda/heap"
	. "go-jvm/src/ch06/utils"
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
	classLoader := heap.NewClassLoader(classPath)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %v\n", cmd.class)
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
