package control

import (
	"MyJavaGolangBook/ch06/instructions/base"
	"MyJavaGolangBook/ch06/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
