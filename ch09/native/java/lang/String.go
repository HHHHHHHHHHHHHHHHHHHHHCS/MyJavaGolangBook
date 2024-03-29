package lang

import (
	"MyJavaGolangBook/ch09/native"
	"MyJavaGolangBook/ch09/rtda"
	"MyJavaGolangBook/ch09/rtda/heap"
)

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
