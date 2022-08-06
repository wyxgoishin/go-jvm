package main

import (
	"fmt"
	"go-jvm/src/ch11/classpath"
	"go-jvm/src/ch11/instruction/base"
	"go-jvm/src/ch11/rtda"
	"go-jvm/src/ch11/rtda/heap"
	"strings"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.classPath)
	classLoader := heap.NewClassLoader(cp, cmd.verbose)
	mainThread := rtda.NewThread()
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  mainThread,
	}
}

func (jvm *JVM) start() {
	jvm.initVM()
	jvm.execMain()
}

func (jvm *JVM) initVM() {
	vmClass := jvm.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(jvm.mainThread, vmClass)
	interpret(jvm.mainThread, jvm.cmd.verbose)
}

func (jvm *JVM) execMain() {
	className := strings.Replace(jvm.cmd.class, ".", "/", -1)
	mainClass := jvm.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("main method not found in class %v\n", jvm.cmd.class)
		return
	}
	argsArr := jvm.createArgsArray()
	frame := jvm.mainThread.NewFrame(mainMethod)
	frame.LocalVars.SetRef(0, argsArr)
	jvm.mainThread.PushFrame(frame)
	interpret(jvm.mainThread, jvm.cmd.verbose)
}

func (jvm *JVM) createArgsArray() *heap.Object {
	strClass := jvm.classLoader.LoadClass("java/lang/String")
	argsArr := strClass.ArrayClass().NewArray(uint(len(jvm.cmd.args)))
	jvArgs := argsArr.Refs()
	for i, arg := range jvm.cmd.args {
		jvArgs[i] = heap.JvString(jvm.classLoader, arg)
	}
	return argsArr
}
