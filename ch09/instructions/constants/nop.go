package constants

import (
	"MyJavaGolangBook/ch09/instructions/base"
	"MyJavaGolangBook/ch09/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
