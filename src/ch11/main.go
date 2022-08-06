package main

import (
	"fmt"
	. "go-jvm/src/ch11/utils"
)

func startJVM(cmd *Cmd) {
	jvm := newJVM(cmd)
	jvm.start()
}

func main() {
	cmd := ParseCmd()
	if cmd.versionFlag {
		fmt.Printf("version %s\n", VERSION)
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
