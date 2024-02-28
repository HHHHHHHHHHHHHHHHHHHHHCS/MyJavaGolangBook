package control

import "MyJavaGolangBook/ch10/instructions/base"
import "MyJavaGolangBook/ch10/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
