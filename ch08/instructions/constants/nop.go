package constants

import (
	"MyJavaGolangBook/ch08/instructions/base"
	"MyJavaGolangBook/ch08/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
