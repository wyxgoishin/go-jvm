package main

import (
	"fmt"
	"go-jvm/src/ch07/instruction"
	"go-jvm/src/ch07/rtda"
	"go-jvm/src/ch07/rtda/heap"
	"go-jvm/src/ch07/utils"
)

func interpret(method *heap.Method, verbose bool) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Printf("LocalVars: %v\n", frame.LocalVars)
			fmt.Printf("OperandStack: %v\n", frame.OperandStack)
		}
	}()
	loop(thread, method.Code(), verbose)
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
