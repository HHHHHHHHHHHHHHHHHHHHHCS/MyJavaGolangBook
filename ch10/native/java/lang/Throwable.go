package lang

import (
	"MyJavaGolangBook/ch10/native"
	"MyJavaGolangBook/ch10/rtda"
)

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {

}
