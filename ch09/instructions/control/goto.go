package control

import (
	"MyJavaGolangBook/ch09/instructions/base"
	"MyJavaGolangBook/ch09/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
