package main

import "fmt"
import "strings"
import "MyJavaGolangBook/ch10/classpath"
import "MyJavaGolangBook/ch10/rtda/heap"

// -Xjre "D:\java-se-8u43-ri\jre" java_test.ParseIntTest 123
func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.10")
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
