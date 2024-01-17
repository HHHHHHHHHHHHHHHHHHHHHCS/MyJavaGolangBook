package constants

import (
	"MyJavaGolangBook/ch10/instructions/base"
	"MyJavaGolangBook/ch10/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
