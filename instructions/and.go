package instructions
import "rt"

type IAND struct{ NoOperandsInstruction }
type LAND struct{ NoOperandsInstruction }


func (self *IAND) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}
	