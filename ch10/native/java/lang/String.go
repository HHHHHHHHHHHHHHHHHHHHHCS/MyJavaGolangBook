package lang

import (
	"MyJavaGolangBook/ch10/native"
	"MyJavaGolangBook/ch10/rtda"
	"MyJavaGolangBook/ch10/rtda/heap"
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
