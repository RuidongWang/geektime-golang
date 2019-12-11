package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	var (
		a int = 1
		b int = 1
	)

	t.Log(a)
	for i := 0; i < 10; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println("hello")
}
