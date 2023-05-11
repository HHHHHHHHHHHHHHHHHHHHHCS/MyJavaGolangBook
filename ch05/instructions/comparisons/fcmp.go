package comparisons

import (
	"MyJavaGolangBook/ch05/instructions/base"
	"MyJavaGolangBook/ch05/rtda"
)

type FCMPG struct {
	base.NoOperandsInstruction
}

type FCMPL struct {
	base.NoOperandsInstruction
}

// 浮点数还存在nan 无法比较
// 当其中一个结果是Nan的时候 fcmpg的结果是1  而 fcmpl 的结果是-1
func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
