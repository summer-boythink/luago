package state

import (
	"luago/ch07/luago/binchunk"
)

type luaState struct {
	stack *luaStack
	proto *binchunk.Prototype
	pc    int
}

func New(stackSize int, proto *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(30),
		proto: proto,
		pc:    0,
	}
}
