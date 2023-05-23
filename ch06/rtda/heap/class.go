package heap

import "MyJavaGolangBook/ch06/classfile"

type Class struct {
	accessFlags        uint16
	name               string // this class name
	superClassName     string
	interfaceNames     []string
	constantPool       *classfile.ConstantPool
	fields             []*Field
	methods            []*Method
	loader             *ClsssLoader
	superClass         *Class
	interfaces         []*Class
	instancesSlotCount uint
	staticSlotCount    uint
	staticVars         *Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	//TODO:
}
