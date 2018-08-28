package mod1

import "github.com/go-stack/stack"

// Mod1 returns a stack.Call from a v1 module.
func Mod1() stack.Call {
	return stack.Caller(0)
}
