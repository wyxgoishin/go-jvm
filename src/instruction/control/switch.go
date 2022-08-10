package control

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
	"go-jvm/src/utils"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (inst *TABLE_SWITCH) FetchOperands(reader *utils.ByteCodeReader) {
	reader.SkipPadding()
	inst.defaultOffset = reader.ReadInt32()
	inst.low = reader.ReadInt32()
	inst.high = reader.ReadInt32()

	jumpOffsetCount := inst.high - inst.low + 1
	inst.jumpOffsets = make([]int32, jumpOffsetCount)
	for idx := range inst.jumpOffsets {
		inst.jumpOffsets[idx] = int32(reader.ReadUint32())
	}
}

func (inst *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= inst.low && index <= inst.high {
		offset = int(inst.jumpOffsets[index-inst.low])
	} else {
		offset = int(inst.defaultOffset)
	}
	base.Branch(frame, offset)
}

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (inst *LOOKUP_SWITCH) FetchOperands(reader *utils.ByteCodeReader) {
	reader.SkipPadding()
	inst.defaultOffset = reader.ReadInt32()
	inst.npairs = reader.ReadInt32()

	inst.matchOffsets = make([]int32, inst.npairs*2)
	for idx := range inst.matchOffsets {
		inst.matchOffsets[idx] = int32(reader.ReadUint32())
	}
}

func (inst *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < inst.npairs*2; i += 2 {
		if inst.matchOffsets[i] == key {
			offset := inst.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(inst.defaultOffset))
}
