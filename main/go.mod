module github.com/go-stack/stack-test/main

go 1.15

require (
	github.com/go-stack/stack v1.8.0
	github.com/go-stack/stack-test/mod1 v1.0.0
	github.com/go-stack/stack-test/mod2/v2 v2.0.0
)

// replace github.com/go-stack/stack-test/mod1 => ../mod1

// replace github.com/go-stack/stack => ../../stack
