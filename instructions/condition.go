package instructions

import "rt"

type IFEQ struct{ BranchInstruction }
type IFNE struct{ BranchInstruction }
type IFLT struct{ BranchInstruction }
type IFLE struct{ BranchInstruction }
type IFGT struct{ BranchInstruction }
type IFGE struct{ BranchInstruction }

type IF_ICMPEQ struct{ BranchInstruction }
type IF_ICMPNE struct{ BranchInstruction }
type IF_ICMPLT struct{ BranchInstruction }
type IF_ICMPLE struct{ BranchInstruction }
type IF_ICMPGT struct{ BranchInstruction }
type IF_ICMPGE struct{ BranchInstruction }

type IF_ACMPEQ struct{ BranchInstruction }
type IF_ACMPNE struct{ BranchInstruction }

type IFNULL struct { BranchInstruction}
type IFNONNULL struct { BranchInstruction }

func (self *IFEQ) Execute(frame *rt.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		Branch(frame, self.Offset)
	}
}


func (self *IF_ICMPNE) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		Branch(frame, self.Offset)
	}
}

func (self *IF_ACMPEQ) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		Branch(frame, self.Offset)
	}
}
	
func (self *IFNULL) Execute(frame *rt.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		Branch(frame, self.Offset)
	}
}


