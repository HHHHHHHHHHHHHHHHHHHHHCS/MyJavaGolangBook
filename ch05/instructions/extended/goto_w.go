package extended

import (
	"MyJavaGolangBook/ch05/instructions/base"
	"MyJavaGolangBook/ch05/rtda"
)

type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
