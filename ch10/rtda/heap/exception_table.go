package heap

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	startPc   int
	endPc     int
	handler   int
	catchType *ClassRef
}

//TODO:
