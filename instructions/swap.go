package instructions
import "rt"

type SWAP struct { NoOperandsInstruction } 

func (self *SWAP) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
	