package main

import (
	"MyJavaGolangBook/ch09/classpath"
	"MyJavaGolangBook/ch09/rtda/heap"
	"fmt"
)
import "strings"

// -Xjre "D:\java-se-8u43-ri\jre" java_test.BubbleSortTest
// -Xjre "D:\java-se-8u43-ri\jre" java_test.HelloWorld chi fan 吃饭, 饿了!
func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.9")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseInstFlag, cmd.args)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
