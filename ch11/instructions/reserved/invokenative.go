package reserved

import (
	"MyJavaGolangBook/ch11/instructions/base"
	"MyJavaGolangBook/ch11/native"
	_ "MyJavaGolangBook/ch11/native/java/lang"
	_ "MyJavaGolangBook/ch11/native/sun/misc"
	"MyJavaGolangBook/ch11/rtda"
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
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
