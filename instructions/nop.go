package instructions

import "rt"

type NOP struct { NoOperandsInstruction }

func (self *NOP) Execute (frame *rt.Frame) {
	// do nothing
}