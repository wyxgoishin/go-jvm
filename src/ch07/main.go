package main

import (
	"flag"
	"fmt"
	"go-jvm/src/ch07/classpath"
	"go-jvm/src/ch07/rtda/heap"
	. "go-jvm/src/ch07/utils"
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
	verbose     bool
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
	flag.BoolVar(&cmd.verbose, "verbose", false, "whether to log")

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
	classLoader := heap.NewClassLoader(classPath, cmd.verbose)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, cmd.verbose)
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
