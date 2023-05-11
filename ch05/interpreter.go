package main

import (
	"MyJavaGolangBook/ch05/classfile"
	"MyJavaGolangBook/ch05/instructions"
	"MyJavaGolangBook/ch05/instructions/base"
	"MyJavaGolangBook/ch05/rtda"
	"fmt"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLoccals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLoccals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {

		fmt.Printf("LocalVars : %v\n", frame.LocalVars())
		fmt.Printf("OperandStack : %v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		// 暂时还没有做static 178 B2
		if opcode == 178 || opcode == 182 {
			frame.SetNextPC(reader.PC())
			continue
		}
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}

}
