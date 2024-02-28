package control

import "MyJavaGolangBook/ch11/instructions/base"
import "MyJavaGolangBook/ch11/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
