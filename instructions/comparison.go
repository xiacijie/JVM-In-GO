package instructions
import "rt"

type LCMP struct{ NoOperandsInstruction }


func (self *LCMP) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
		} else if v1 == v2 {
			stack.PushInt(0)
		} else {
			stack.PushInt(-1)
	}
}
	

type FCMPG struct{ NoOperandsInstruction }
type FCMPL struct{ NoOperandsInstruction }

func _fcmp(frame *rt.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (self *FCMPG) Execute(frame *rt.Frame) {
	_fcmp(frame, true)
}

func (self *FCMPL) Execute(frame *rt.Frame) {
	_fcmp(frame, false)
}