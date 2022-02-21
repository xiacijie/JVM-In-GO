package instructions

import "rt"

type POP struct{ NoOperandsInstruction }
type POP2 struct{ NoOperandsInstruction }

func (self *POP) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (self *POP2) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
	