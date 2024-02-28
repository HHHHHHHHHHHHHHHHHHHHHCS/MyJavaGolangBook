package main

import (
	"fmt"
)

// -Xjre "D:\java-se-8u43-ri\jre" java_test.ParseIntTest 123
func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.11")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		newJVM(cmd).start()
	}
}
