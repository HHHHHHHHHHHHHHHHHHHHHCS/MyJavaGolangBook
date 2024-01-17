package control

import (
	"MyJavaGolangBook/ch10/instructions/base"
	"MyJavaGolangBook/ch10/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
