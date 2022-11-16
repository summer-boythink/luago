package state

import . "luago/ch09/luago/api"

import "luago/ch09/luago/binchunk"

type closure struct {
	proto *binchunk.Prototype //lua closure
	goFunc GoFunction 
}

func newLuaClosure(proto *binchunk.Prototype) *closure {
	return &closure{proto: proto}
}

func newGoClosure(f GoFunction) *closure {
	return &closure{goFunc: f}
}
