package instructions
import "rt"

type GOTO struct{ BranchInstruction }

func (self *GOTO) Execute(frame *rt.Frame) {
	Branch(frame, self.Offset)
}

type TABLE_SWITCH struct {
	defaultOffset int32 
	low int32 
	high int32 
	jumpOffset []int32 
}

func (self *TABLE_SWITCH) FetchOperands(reader* BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset =  reader.ReadInt32()
	self.low =  reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffset = reader.ReadInt32s(jumpOffsetsCount)
}

func (self* TABLE_SWITCH) Execute(frame *rt.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int 
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffset[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}

	Branch(frame, offset)
}
	
type LOOPUP_SWITCH struct {
	defaultOffset int32 
	npairs int32 
	matchOffetests []int32 
}

func (self* LOOPUP_SWITCH) FetchOperands(reader* BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffetests = reader.ReadInt32s(self.npairs * 2)
}

func (self* LOOPUP_SWITCH) Execute(frame* rt.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffetests[i] == key {
			offset := self.matchOffetests[i+1]
			Branch(frame, int(offset))
			return  
		}
	}
	Branch(frame, int(self.defaultOffset))
}