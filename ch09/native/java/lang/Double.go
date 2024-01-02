package lang

import (
	"MyJavaGolangBook/ch09/native"
	"MyJavaGolangBook/ch09/rtda"
	"math"
)

const jlDouble = "java/lang/Double"

func init() {
	native.Register(jlDouble, "doubleToRawIntBits", "(D)J", doubleToRawIntBits)
	native.Register(jlDouble, "longBitsToDouble", "(J)D", longBitsToDouble)

}

func doubleToRawIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OperandStack().PushLong(int64(bits))
}

func longBitsToDouble(frame *rtda.Frame) {
	bits := frame.LocalVars().GetLong(0)
	value := math.Float64frombits(uint64(bits))
	frame.OperandStack().PushDouble(value)
}
