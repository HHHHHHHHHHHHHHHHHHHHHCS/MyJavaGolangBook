package control

import (
	"MyJavaGolangBook/ch08/instructions/base"
	"MyJavaGolangBook/ch08/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
