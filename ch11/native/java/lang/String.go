package lang

import (
	"MyJavaGolangBook/ch11/native"
	"MyJavaGolangBook/ch11/rtda"
	"MyJavaGolangBook/ch11/rtda/heap"
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
