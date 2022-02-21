package instructions
import "rt"

type WIDE struct {
	modifiedInstruction Instruction
}

func (self *WIDE) FetchOperands(reader* BytecodeReader) {
	opcode := reader.ReadInt8()
	switch opcode {
		case 0x15:  // iload
			inst := &ILOAD{}
			inst.Index = uint(reader.ReadInt16())
			self.modifiedInstruction = inst 
		case 0x16: ... // lload
		case 0x17: ... // fload
		case 0x18: ... // dload
		case 0x19: ... // aload
		case 0x36: ... // istore
		case 0x37: ... // lstore
		case 0x38: ... // fstore
		case 0x39: ... // dstore
		case 0x3a: ... // astore
		case 0x84:  // iinc
			inst := IINC{} 
			inst.Index = uint(reader.ReadUint16())
			inst.Const = int32(reader.ReadInt16())
			self.modifiedInstruction = inst 
		case 0xa9: // ret
		panic("Unsupported opcode: 0xa9!")
	}
}


func (self *WIDE) Execute(frame *rt.Frame) {
	self.modiself.modifiedInstruction.Execute(frame)
}

type GOTO_W struct {
	offset int 
}

func (self *GOTO_W) FetchOperands(reader* BytecodeReader) {
	self.offset =  int(reader.ReadInt32())

}

func (self *GOTO_W) Execute(frame *rt.Frame) {
	Branch(frame, self.offset)
}

