package constants

import "MyJavaGolangBook/ch11/instructions/base"
import "MyJavaGolangBook/ch11/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
