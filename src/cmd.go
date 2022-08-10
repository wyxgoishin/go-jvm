package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	classPath   string
	XjreOption  string
	class       string
	args        []string
	verbose     bool
}

func ParseCmd() *Cmd {
	cmd := new(Cmd)

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
