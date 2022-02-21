package instructions
import "rt"

type ISHL struct{ NoOperandsInstruction } 
type ISHR struct{ NoOperandsInstruction } 
type IUSHR struct{ NoOperandsInstruction } 
type LSHL struct{ NoOperandsInstruction } 
type LSHR struct{ NoOperandsInstruction }
type LUSHR struct{ NoOperandsInstruction } 


func (self *ISHL) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}


func (self *LSHR) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

func (self *IUSHR) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}
	