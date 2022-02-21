package instructions
import "math"
import "rt"

type DREM struct{ NoOperandsInstruction }
type FREM struct{ NoOperandsInstruction }
type IREM struct{ NoOperandsInstruction }
type LREM struct{ NoOperandsInstruction }

func (self *IREM) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

func (self *DREM) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}