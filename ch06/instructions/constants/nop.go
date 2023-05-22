package constants

import (
	"MyJavaGolangBook/ch06/instructions/base"
	"MyJavaGolangBook/ch06/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
