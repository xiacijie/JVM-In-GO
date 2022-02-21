package instructions
import "rt"

type D2F struct{ NoOperandsInstruction }
type D2I struct{ NoOperandsInstruction }
type D2L struct{ NoOperandsInstruction }


func (self *D2I) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}
	