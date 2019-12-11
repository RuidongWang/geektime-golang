package constant_test

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wendnesday
	Friday
)

const (
	readable = 1 << iota
	writeable
	excutable
)

func TestConstantTry(t *testing.T) {
	t.Log("this is the log", Monday, Tuesday)
	fmt.Println("this is print", Monday, Tuesday)
}
