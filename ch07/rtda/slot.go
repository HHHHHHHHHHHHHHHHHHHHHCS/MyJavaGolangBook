package rtda

import "MyJavaGolangBook/ch07/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
