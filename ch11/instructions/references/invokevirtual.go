package references

import (
	"MyJavaGolangBook/ch11/instructions/base"
	"MyJavaGolangBook/ch11/rtda"
	"MyJavaGolangBook/ch11/rtda/heap"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {

		// hack!
		if methodRef.Name() == "digit" {
			_digit(frame.OperandStack(), methodRef.Descriptor())
			return
		}

		panic("java.lang.NullPointerException")
	}

	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		if !(ref.Class().IsArray() && resolvedMethod.Name() == "clone") {
			panic("java.lang.IllegalAccessError")
		}
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

// hack!
func _digit(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(II)I":
		stack.PopInt() //radix 默认十进制不用管
		stack.PopInt() //char1
		char0 := stack.PopInt()
		if char0 < '0' || char0 > '9' {
			panic("digit fail!")
		} else {
			stack.PushInt(char0 - '0')
		}
	}
}
