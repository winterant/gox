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
		fmt.Printf("%v\n", err)
		fmt.Printf("%+v\n", err) // including call stack
		fmt.Println(xerrors.Cause(err))
	}

	err = c()
	if err != nil {
		fmt.Println("--------------- go std error ---------------")
		fmt.Printf("%v\n", err)
		fmt.Printf("%+v\n", err) // including call stack
		fmt.Println(xerrors.Cause(err))
	}
}
