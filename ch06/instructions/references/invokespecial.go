package references

import (
	"MyJavaGolangBook/ch06/instructions/base"
	"MyJavaGolangBook/ch06/rtda"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
