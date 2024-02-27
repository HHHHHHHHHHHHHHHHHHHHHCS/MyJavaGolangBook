package control

import (
	"MyJavaGolangBook/ch11/instructions/base"
	"MyJavaGolangBook/ch11/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
