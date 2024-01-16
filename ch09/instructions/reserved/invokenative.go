package reserved

import (
	"MyJavaGolangBook/ch09/instructions/base"
	"MyJavaGolangBook/ch09/native"
	_ "MyJavaGolangBook/ch09/native/java/lang"
	_ "MyJavaGolangBook/ch09/native/sun/misc"
	"MyJavaGolangBook/ch09/rtda"
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
