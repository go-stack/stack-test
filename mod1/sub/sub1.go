package sub1diff

import "github.com/go-stack/stack"

// Sub returns a stack.Call from a v1 module.
func Sub() stack.Call {
	return stack.Caller(0)
}
