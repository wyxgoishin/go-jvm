package control

import (
	"go-jvm/src/ch07/instruction/base"
	"go-jvm/src/ch07/rtda"
	"go-jvm/src/ch07/utils"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (t *TABLE_SWITCH) FetchOperands(cr *utils.ByteCodeReader) {
	cr.SkipPadding()
	t.defaultOffset = int32(cr.ReadUint32())
	t.low = int32(cr.ReadUint32())
	t.high = int32(cr.ReadUint32())

	jumpOffsetCount := t.high - t.low + 1
	t.jumpOffsets = make([]int32, jumpOffsetCount)
	for idx := range t.jumpOffsets {
		t.jumpOffsets[idx] = int32(cr.ReadUint32())
	}
}

func (t *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack.PopInt()
	var offset int
	if index >= t.low && index <= t.high {
		offset = int(t.jumpOffsets[index-t.low])
	} else {
		offset = int(t.defaultOffset)
	}
	base.Branch(frame, offset)
}

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (l *LOOKUP_SWITCH) FetchOperands(cr *utils.ByteCodeReader) {
	cr.SkipPadding()
	l.defaultOffset = int32(cr.ReadUint32())
	l.npairs = int32(cr.ReadUint32())

	l.matchOffsets = make([]int32, l.npairs*2)
	for idx := range l.matchOffsets {
		l.matchOffsets[idx] = int32(cr.ReadUint32())
	}
}

func (l *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack.PopInt()
	for i := int32(0); i < l.npairs*2; i += 2 {
		if l.matchOffsets[i] == key {
			offset := l.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(l.defaultOffset))
}
