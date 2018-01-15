package test

import (
	"testing"
	"fmt"
)

func TestFoo(t *testing.T) {
	r, e :=Foo(5)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println("result:", r)
}