package lang

import (
	"MyJavaGolangBook/ch09/native"
	"MyJavaGolangBook/ch09/rtda"
	"math"
)

const jlFloat = "java/lang/Float"

func init() {
	native.Register(jlFloat, "floatToRawIntBits", "(F)I", floatToRawIntBits)
	native.Register(jlFloat, "intBitsToFloat", "(I)F", intBitsToFloat)
}

func floatToRawIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}

func intBitsToFloat(frame *rtda.Frame) {
	bits := frame.LocalVars().GetInt(0)
	value := math.Float32frombits(uint32(bits))
	frame.OperandStack().PushFloat(value)
}
