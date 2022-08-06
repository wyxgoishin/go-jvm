package main

import (
	"fmt"
	"go-jvm/src/ch11/instruction"
	"go-jvm/src/ch11/rtda"
	"go-jvm/src/ch11/utils"
)

func interpret(thread *rtda.Thread, verbose bool) {
	loop(thread, verbose)
}

func loop(thread *rtda.Thread, verbose bool) {
	cr := &utils.ByteCodeReader{}
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
