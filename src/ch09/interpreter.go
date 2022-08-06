package main

import (
	"fmt"
	"go-jvm/src/ch09/instruction"
	"go-jvm/src/ch09/rtda"
	"go-jvm/src/ch09/rtda/heap"
	"go-jvm/src/ch09/utils"
)

func interpret(method *heap.Method, verbose bool, args []string) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	jvArgs := createArgsArray(method.Class().ClassLoader(), args)
	frame.LocalVars.SetRef(0, jvArgs)
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println(r)
	//		fmt.Printf("LocalVars: %v\n", frame.LocalVars)
	//		fmt.Printf("OperandStack: %v\n", frame.OperandStack)
	//	}
	//}()
	loop(thread, method.Code(), verbose)
}

func createArgsArray(classloader *heap.ClassLoader, args []string) *heap.Object {
	strClass := classloader.LoadClass("java/lang/String")
	argsArr := strClass.ArrayClass().NewArray(uint(len(args)))
	jvArgs := argsArr.Refs()
	for i, arg := range args {
		jvArgs[i] = heap.JvString(classloader, arg)
	}
	return argsArr
}

func loop(thread *rtda.Thread, bytecode []byte, verbose bool) {
	cr := utils.NewByteCodeReader(bytecode)
	for {
		frame := thread.CurrentFrame()
		if frame == nil {
			break
		}
		pc := frame.NextPC()
		thread.SetPC(pc)
		cr.Reset(frame.Method().Code(), pc)
		opcode := cr.ReadUint8()
		inst := instruction.NewInstruction(opcode)
		inst.FetchOperands(cr)
		frame.SetNextPC(cr.PC())
		inst.Execute(frame)
		if verbose {
			fmt.Printf("pc: %d; inst: %T %v; %v.%v()\n", pc, inst, inst, frame.Method().Class().Name(), frame.Method().Name())
		}
	}
}
