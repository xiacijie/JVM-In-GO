package classfile
type ConstantPool []ConstantInfo 

func readConstantPool(reader* ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
}