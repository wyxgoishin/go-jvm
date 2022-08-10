package extend

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/instruction/load"
	"go-jvm/src/instruction/math"
	"go-jvm/src/instruction/store"
	"go-jvm/src/rtda"
	"go-jvm/src/utils"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (inst *WIDE) FetchOperands(reader *utils.ByteCodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		instruction := new(load.ILOAD)
		instruction.Index = uint(reader.ReadUint16())
		inst.modifiedInstruction = instruction
	case 0x16:
		instruction := new(load.LLOAD)
		instruction.Index = uint(reader.ReadUint16())
		inst.modifiedInstruction = instruction
	case 0x17:
		instruction := new(load.FLOAD)
		instruction.Index = uint(reader.ReadUint16())
		inst.modifiedInstruction = instruction
	case 0x18:
		instruction := new(load.DLOAD)
		instruction.Index = uint(reader.ReadUint16())
		inst.modifiedInstruction = instruction
	case 0x19:
		instruction := new(load.ALOAD)
		instruction.Index = uint(reader.ReadUint16())
		inst.modifiedInstruction = instruction
	case 0x36:
		instruction := new(store.ISTORE)
		instruction.Index = uint(reader.ReadUint16())
		inst.modifiedInstruction = instruction
	case 0x37:
		instruction := new(store.LSTORE)
		instruction.Index = uint(reader.ReadUint16())
		inst.modifiedInstruction = instruction
	case 0x38:
		instruction := new(store.FSTORE)
		instruction.Index = uint(reader.ReadUint16())
		inst.modifiedInstruction = instruction
	case 0x39:
		instruction := new(store.DSTORE)
		instruction.Index = uint(reader.ReadUint16())
		inst.modifiedInstruction = instruction
	case 0x3a:
		instruction := new(store.ASTORE)
		instruction.Index = uint(reader.ReadUint16())
		inst.modifiedInstruction = instruction
	case 0x84:
		instruction := new(math.IINC)
		instruction.Index = uint(reader.ReadUint16())
		instruction.Const = int32(reader.ReadUint16())
		inst.modifiedInstruction = instruction
	case 0xa9:
		panic("ret waiting to be implemented")
	}
}

func (inst *WIDE) Execute(frame *rtda.Frame) {
	inst.modifiedInstruction.Execute(frame)
}
