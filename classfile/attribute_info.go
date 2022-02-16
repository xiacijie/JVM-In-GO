package classfile

type AttributeInfo interface {
	readerInfo(reader* ClassReader)
}

func readerAttributes(reader* ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfom , attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}

	return attributes 
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}