package extend

import (
	"go-jvm/src/ch07/instruction/base"
	"go-jvm/src/ch07/instruction/load"
	"go-jvm/src/ch07/instruction/math"
	"go-jvm/src/ch07/instruction/store"
	"go-jvm/src/ch07/rtda"
	"go-jvm/src/ch07/utils"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (w *WIDE) FetchOperands(cr *utils.ByteCodeReader) {
	opcode := cr.ReadUint8()
	switch opcode {
	case 0x15:
		instruction := new(load.ILOAD)
		instruction.Index = uint(cr.ReadUint16())
		w.modifiedInstruction = instruction
	case 0x16:
		instruction := new(load.LLOAD)
		instruction.Index = uint(cr.ReadUint16())
		w.modifiedInstruction = instruction
	case 0x17:
		instruction := new(load.FLOAD)
		instruction.Index = uint(cr.ReadUint16())
		w.modifiedInstruction = instruction
	case 0x18:
		instruction := new(load.DLOAD)
		instruction.Index = uint(cr.ReadUint16())
		w.modifiedInstruction = instruction
	case 0x19:
		instruction := new(load.ALOAD)
		instruction.Index = uint(cr.ReadUint16())
		w.modifiedInstruction = instruction
	case 0x36:
		instruction := new(store.ISTORE)
		instruction.Index = uint(cr.ReadUint16())
		w.modifiedInstruction = instruction
	case 0x37:
		instruction := new(store.LSTORE)
		instruction.Index = uint(cr.ReadUint16())
		w.modifiedInstruction = instruction
	case 0x38:
		instruction := new(store.FSTORE)
		instruction.Index = uint(cr.ReadUint16())
		w.modifiedInstruction = instruction
	case 0x39:
		instruction := new(store.DSTORE)
		instruction.Index = uint(cr.ReadUint16())
		w.modifiedInstruction = instruction
	case 0x3a:
		instruction := new(store.ASTORE)
		instruction.Index = uint(cr.ReadUint16())
		w.modifiedInstruction = instruction
	case 0x84:
		instruction := new(math.IINC)
		instruction.Index = uint(cr.ReadUint16())
		instruction.Const = int32(cr.ReadUint16())
		w.modifiedInstruction = instruction
	case 0xa9:
		panic("ret waiting to be implemented")
	}
}

func (w *WIDE) Execute(frame *rtda.Frame) {
	w.modifiedInstruction.Execute(frame)
}
