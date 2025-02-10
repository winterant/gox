package main

import (
	"errors"
	"fmt"

	"github.com/winterant/gox/pkg/xerrors"
)

func xa() error {
	return xerrors.New("xerrors new error")
}

func xb() error {
	err := xa()
	if err != nil {
		return xerrors.Wrap(err, "xerrors wrap error in b()")
	}
	return nil
}

func xc() error {
	err := xb()
	if err != nil {
		return xerrors.Wrap(err, "xerrors wrap error in c()")
	}
	return nil
}

func a() error {
	return errors.New("go std error")
}

func b() error {
	err := a()
	if err != nil {
		return xerrors.Wrap(err, "xerrors wrap error in b()")
	}
	return nil
}

func c() error {
	err := b()
	if err != nil {
		return xerrors.Wrap(err, "xerrors wrap error in c()")
	}
	return nil
}

func main() {
	err := xc()
	if err != nil {
		fmt.Println("--------------- xerrors ---------------")
		fmt.Printf("[normal]: %v\n", err)
		fmt.Printf("[stack]: %+v\n", err) // including call stack
		fmt.Println("[Cause]:", xerrors.Cause(err))
	}

	err = c()
	if err != nil {
		fmt.Println("--------------- go std error ---------------")
		fmt.Printf("[normal]: %v\n", err)
		fmt.Printf("[stack]: %+v\n", err) // including call stack
		fmt.Println("[Cause]:", xerrors.Cause(err))
	}
}

/* Output:
--------------- xerrors ---------------
[normal]: xerrors wrap error in c(): xerrors wrap error in b(): xerrors new error
[stack]: xerrors wrap error in c(): xerrors wrap error in b(): xerrors new error
        /Users/zhaojinglong01/Personal/Projects/gox/examples/xerrors/main.go:11 main.xa
        /Users/zhaojinglong01/Personal/Projects/gox/examples/xerrors/main.go:15 main.xb
        /Users/zhaojinglong01/Personal/Projects/gox/examples/xerrors/main.go:23 main.xc
        /Users/zhaojinglong01/Personal/Projects/gox/examples/xerrors/main.go:51 main.main
        /Library/env/goenv/versions/1.21.13/src/runtime/proc.go:267 runtime.main
        /Library/env/goenv/versions/1.21.13/src/runtime/asm_arm64.s:1197 runtime.goexit
[Cause]: xerrors new error
--------------- go std error ---------------
[normal]: xerrors wrap error in c(): xerrors wrap error in b(): go std error
[stack]: xerrors wrap error in c(): xerrors wrap error in b(): go std error
        /Users/zhaojinglong01/Personal/Projects/gox/examples/xerrors/main.go:37 main.b
        /Users/zhaojinglong01/Personal/Projects/gox/examples/xerrors/main.go:43 main.c
        /Users/zhaojinglong01/Personal/Projects/gox/examples/xerrors/main.go:59 main.main
        /Library/env/goenv/versions/1.21.13/src/runtime/proc.go:267 runtime.main
        /Library/env/goenv/versions/1.21.13/src/runtime/asm_arm64.s:1197 runtime.goexit
[Cause]: go std error
*/
