package main

import (
	"fmt"
	"go-jvm/src/ch06/instruction"
	"go-jvm/src/ch06/rtda"
	"go-jvm/src/ch06/rtda/heap"
	"go-jvm/src/ch06/utils"
)

func interpret(method *heap.Method) {
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
	loop(thread, method.Code())
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	cr := utils.NewByteCodeReader(bytecode)
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		cr.SetPC(pc)
		opcode := cr.ReadUint8()
		inst := instruction.NewInstruction(opcode)
		inst.FetchOperands(cr)
		frame.SetNextPC(cr.PC())
		inst.Execute(frame)
		fmt.Printf("pc: %d; inst: %T %v\n", pc, inst, inst)
	}
}
