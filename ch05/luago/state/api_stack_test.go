package state

import (
	"reflect"
	"testing"
)

var l *luaState = &luaState{
	&luaStack{
		slots: []luaValue{'a', 'b', 'c', 'd', 'e'},
		top:   5,
	},
}

func TestRotate(t *testing.T) {
	l.Rotate(2, 1)
	if !reflect.DeepEqual(l.stack.slots, []luaValue{'a', 'e', 'b', 'c', 'd'}) {
		t.Error("err")
	}
}
