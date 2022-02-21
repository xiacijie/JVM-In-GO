package instructions
import "rt"

type IINC struct {
	Index uint
	Const int32
}


func (self *IINC) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}
	
func (self *IINC) Execute(frame *rt.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
	
	