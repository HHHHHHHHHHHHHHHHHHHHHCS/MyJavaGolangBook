package rtda

import "MyJavaGolangBook/ch10/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
