package main

import (
	"flag"
	"fmt"
	"os"
)

type cmd struct {
	helpFlag    bool
	versionFlag bool
	classPath   string
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
	fmt.Printf("classpath: %q; class: %q; args: %q\n", cmd.classPath, cmd.class, cmd.args)
}

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version %s\n", VERSION)
	}
	if cmd.helpFlag || cmd.class == "" {
		printUsage()
	}
	startJVM(cmd)
}
