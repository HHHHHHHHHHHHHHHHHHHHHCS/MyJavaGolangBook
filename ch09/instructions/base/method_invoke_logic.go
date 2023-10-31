package base

import (
	"MyJavaGolangBook/ch08/rtda"
	"MyJavaGolangBook/ch08/rtda/heap"
	"fmt"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	// hack: Object是所有类的超类, 会导致奔溃
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method:%v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
