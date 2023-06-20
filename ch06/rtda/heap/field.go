package heap

import (
	"MyJavaGolangBook/ch06/classfile"
)

type Field struct {
	ClassMember
	slotId uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfFields := range cfFields {
		field := Field{}
		field.class = class
		field.copyMemberInfo(cfFields)
		fields[i] = &field
	}
	return fields
}
