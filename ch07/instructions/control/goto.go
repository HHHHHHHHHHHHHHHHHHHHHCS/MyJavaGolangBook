package control

import (
	"MyJavaGolangBook/ch07/instructions/base"
	"MyJavaGolangBook/ch07/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
