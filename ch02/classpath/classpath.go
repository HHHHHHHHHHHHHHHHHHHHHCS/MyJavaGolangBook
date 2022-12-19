package classpath

type ClassPath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *ClassPath {

}

func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {

}

func (self *ClassPath) String() string {

}
