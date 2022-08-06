package main

import (
	"flag"
	"fmt"
	"go-jvm/src/ch02/classpath"
	. "go-jvm/src/ch02/utils"
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
	cp := classpath.Parse(cmd.XjreOption, cmd.classPath)
	fmt.Printf("classpath: %q\nclass: %q\nargs: %q\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("could not find or load main class %q\n", cmd.class)
		return
	}
	fmt.Printf("class data: %q\n", classData)
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
