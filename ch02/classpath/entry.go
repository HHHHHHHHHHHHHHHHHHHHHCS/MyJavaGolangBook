package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	string() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		//todo:
	}
}
