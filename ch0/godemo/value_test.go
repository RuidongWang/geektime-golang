package godemo

import (
	"testing"
)

var a = 10

func TestDemo(t *testing.T) {
	c := 200
	c = a
	t.Log("赋值操作，把 a 赋值给 c，所以 c 的值为：", c)
	c += a
	t.Log("相加和赋值运算符，实际为 c = c + a，所以 c 的值为：", c)
	c -= a
	t.Log("相减和赋值运算符，实际为 c = c - a，所以 c 的值为：", c)
	b := 7
	b = b &^ 11
	t.Log("按位清零操作", b)

	t.Log(a)

	a := "hello"

	t.Log(a)
}
