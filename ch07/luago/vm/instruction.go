package vm

import "luago/ch07/luago/api"

const MAXARG_Bx = 1<<18 - 1       // 262143
const MAXARG_sBx = MAXARG_Bx >> 1 // 131071

/*
 31       22       13       5    0
  +-------+^------+-^-----+-^-----
  |b=9bits |c=9bits |a=8bits|op=6|
  +-------+^------+-^-----+-^-----
  |    bx=18bits    |a=8bits|op=6|
  +-------+^------+-^-----+-^-----
  |   sbx=18bits    |a=8bits|op=6|
  +-------+^------+-^-----+-^-----
  |    ax=26bits            |op=6|
  +-------+^------+-^-----+-^-----
 31      23      15       7      0
*/
type Instruction uint32

func (in Instruction) Opcode() int {
	return int(in & 0x3F)
}

func (in Instruction) ABC() (a, b, c int) {
	a = int(in >> 6 & 0xFF)
	c = int(in >> 14 & 0x1FF)
	b = int(in >> 23 & 0x1FF)
	return
}

func (in Instruction) ABx() (a, bx int) {
	a = int(in >> 6 & 0xFF)
	bx = int(in >> 14)
	return
}

func (in Instruction) AsBx() (a, sbx int) {
	a, bx := in.ABx()
	return a, bx - MAXARG_sBx
}

func (in Instruction) Ax() int {
	return int(in >> 6)
}

func (in Instruction) OpName() string {
	return opcodes[in.Opcode()].name
}

func (in Instruction) OpMode() byte {
	return opcodes[in.Opcode()].opMode

}
func (in Instruction) BMode() byte {
	return opcodes[in.Opcode()].argBMode
}

func (in Instruction) CMode() byte {
	return opcodes[in.Opcode()].argCMode
}

func (s Instruction) Execute(vm api.LuaVM) {
	action := opcodes[s.Opcode()].action
	if action != nil {
		action(s, vm)
	} else {
		panic(s.OpName())
	}
}
