package constants

import (
	"MyJavaGolangBook/ch05/instructions/base"
	"MyJavaGolangBook/ch05/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
