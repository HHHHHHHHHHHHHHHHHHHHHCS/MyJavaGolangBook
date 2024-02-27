package main

import (
	"MyJavaGolangBook/ch11/classpath"
	"MyJavaGolangBook/ch11/instructions/base"
	"MyJavaGolangBook/ch11/rtda"
	"MyJavaGolangBook/ch11/rtda/heap"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.NewThread(),
	}
}
func (self *JVM) start() {
	self.initVM()
	self.execMain()
}

func (self *JVM) initVM() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	interpret(self.mainThread, self.cmd.verboseInstFlag)
}

//TODO:
