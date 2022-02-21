package instructions
import "rt"

type BIPUSH struct { val int8 }
type SIPUSH struct { val int16 }

func (self *BIPUSH) FetchOperands(reader* BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame* rt.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}


