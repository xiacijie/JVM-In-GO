package interpreter

import (
	"classfile"
	"rt"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	thread := rt.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.pushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}
