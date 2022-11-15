package state

type luaState struct {
	stack *luaStack
	// proto *binchunk.Prototype
	// pc    int
}

func New() *luaState {
	return &luaState{
		stack: newLuaStack(30),
		// proto: proto,
		// pc:    0,
	}
}

func (self *luaState) pushLuaStack(stack *luaStack) {
	stack.prev = self.stack
	self.stack = stack
}

func (self *luaState) popLuaStack() {
	stack := self.stack
	self.stack = stack.prev
	stack.prev = nil
}
