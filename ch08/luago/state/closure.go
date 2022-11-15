package state

import "luago/ch08/luago/binchunk"

type closure struct {
	proto *binchunk.Prototype
}

func newLuaClosure(proto *binchunk.Prototype) *closure {
	return &closure{proto: proto}
}