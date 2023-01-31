package classfile

type ConstantPool []ConstantInfo

// 表头大小 = 实际上的常量池大小+1
// 有效常量池索引是1~n-1
// CONSTANT_Long_info 和 CONSTANT_Long_Double_info各占两个位置, 1~n-1的某些数也会变成无效数字

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	// 索引从1开始
	for i := 1; i < cpCount; i++ {
		info = readConstantInfo(reader, cp)
		cp[i] = info
		switch info.(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ //占着两个位置
		}
	}
	return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
