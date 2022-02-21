package instructions
import "rt"

type ILOAD struct{ Index8Instruction }
type ILOAD_0 struct{ NoOperandsInstruction }
type ILOAD_1 struct{ NoOperandsInstruction }
type ILOAD_2 struct{ NoOperandsInstruction }
type ILOAD_3 struct{ NoOperandsInstruction }

func _iload(frame *rt.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (self* ILOAD) Execute (frame *rt.Frame) {
	_iload(frame, uint(self.Index))
}

func (self *ILOAD_1) Execute(frame *rt.Frame) {
	_iload(frame, 1)
}
	