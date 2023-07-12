package main

import (
	"MyJavaGolangBook/ch06/classpath"
	"MyJavaGolangBook/ch06/rtda/heap"
	"fmt"
)
import "strings"

// java_test.MyObject
func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.6")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
