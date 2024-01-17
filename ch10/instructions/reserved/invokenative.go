package reserved

import (
	"MyJavaGolangBook/ch10/instructions/base"
	"MyJavaGolangBook/ch10/native"
	_ "MyJavaGolangBook/ch10/native/java/lang"
	_ "MyJavaGolangBook/ch10/native/sun/misc"
	"MyJavaGolangBook/ch10/rtda"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfieldLinkError: " + methodInfo)
	}
	nativeMethod(frame)
}
