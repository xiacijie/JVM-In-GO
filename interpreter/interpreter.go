package interpreter

import "classfile"
import "intructions"
import "rt"

func interpret(methodInfo * classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

}