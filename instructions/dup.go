package instructions

import "rt"

type DUP struct{ NoOperandsInstruction }
type DUP_X1 struct{ NoOperandsInstruction }
type DUP_X2 struct{ NoOperandsInstruction }
type DUP2 struct{ NoOperandsInstruction }
type DUP2_X1 struct{ NoOperandsInstruction }
type DUP2_X2 struct{ NoOperandsInstruction }


func (self *DUP) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}