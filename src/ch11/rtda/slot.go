package rtda

import "go-jvm/src/ch11/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
