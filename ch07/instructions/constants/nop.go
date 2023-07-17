package constants

import (
	"MyJavaGolangBook/ch07/instructions/base"
	"MyJavaGolangBook/ch07/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
