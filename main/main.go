package main

import (
	"fmt"

	"github.com/go-stack/stack"
	"github.com/go-stack/stack-test/mod1"
	sub1diff "github.com/go-stack/stack-test/mod1/sub"
	"github.com/go-stack/stack-test/mod2/v2"
)

func main() {
	calls := []stack.Call{
		mod1.Mod1(),
		sub1diff.Sub(),
		mod2.Mod2(),
	}

	for _, c := range calls {
		for _, fs := range []string{"%n", "%+n", "%#s", "%+s"} {
			fmt.Printf("%v: "+fs+"\n", fs, c)
		}

		fmt.Println()
	}
}
