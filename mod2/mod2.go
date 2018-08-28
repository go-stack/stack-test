package mod2

import "github.com/go-stack/stack"

// Mod2 returns a stack.Call from a v2 module.
func Mod2() stack.Call {
	return stack.Caller(0)
}
