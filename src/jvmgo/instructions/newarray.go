package instructions

import (
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Create new array
type newarray struct {
    atype uint8
}
func (self *newarray) fetchOperands(bcr *BytecodeReader) {
    self.atype = bcr.readUint8()
}
func (self *newarray) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    count := stack.PopInt()
    classLoader := frame.Method().Class().ClassLoader()
    arr := rtc.NewPrimitiveArray(self.atype, count, classLoader)
    stack.PushRef(arr)
}

// Create new array of reference
type anewarray struct {Index16Instruction}
func (self *anewarray) Execute(frame *rtda.Frame) {
    cp := frame.Method().Class().ConstantPool()
    cClass := cp.GetConstant(self.index).(*rtc.ConstantClass)
    componentClass := cClass.Class()

    if componentClass.InitializationNotStarted() {
        thread := frame.Thread()
        frame.SetNextPC(thread.PC()) // undo anewarray
        initClass(componentClass, thread)
        return
    }

    stack := frame.OperandStack()
    count := stack.PopInt()
    classLoader := frame.Method().Class().ClassLoader()
    arr := rtc.NewRefArray(componentClass, count, classLoader)
    stack.PushRef(arr)
}

// Create new multidimensional array
type multianewarray struct {
    index       uint16
    dimensions  uint8
}
func (self *multianewarray) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
    self.dimensions = bcr.readUint8()
}
func (self *multianewarray) Execute(frame *rtda.Frame) {
    // todo
    panic("todo multianewarray")
}
