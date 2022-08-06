package main

import (
	"fmt"
	"go-jvm/src/ch05/classfile"
	"go-jvm/src/ch05/instruction"
	"go-jvm/src/ch05/rtda"
	"go-jvm/src/ch05/utils"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	byteCode := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	thread.PushFrame(frame)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Printf("LocalVars: %v\n", frame.LocalVars)
			fmt.Printf("OperandStack: %v\n", frame.OperandStack)
		}
	}()
	loop(thread, byteCode)
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	cr := utils.NewByteCodeReader(bytecode)
	for {
		//fmt.Printf("OperandStack: %v, LocalVars: %v\n", frame.OperandStack, frame.LocalVars)
		pc := frame.NextPC()
		thread.SetPC(pc)
		cr.SetPC(pc)
		opcode := cr.ReadUint8()
		inst := instruction.NewInstruction(opcode)
		inst.FetchOperands(cr)
		frame.SetNextPC(cr.PC())
		inst.Execute(frame)
		fmt.Printf("pc: %d; inst: %T %v\n", pc, inst, inst)
		fmt.Printf("LocalVars: %v; OperandStack: %v\n", frame.LocalVars, frame.OperandStack)
	}
}
