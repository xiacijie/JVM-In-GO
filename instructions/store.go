package instructions

import "rt"

type LSTORE struct{ Index8Instruction }
type LSTORE_0 struct{ NoOperandsInstruction }
type LSTORE_1 struct{ NoOperandsInstruction }
type LSTORE_2 struct{ NoOperandsInstruction }
type LSTORE_3 struct{ NoOperandsInstruction }

func _lstore(frame *rt.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (self *LSTORE) Execute(frame *rt.Frame) {
	_lstore(frame, uint(self.Index))
}

func (self *LSTORE_2) Execute(frame *rt.Frame) {
	_lstore(frame, 2)
}
	
	