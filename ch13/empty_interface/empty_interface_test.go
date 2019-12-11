package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("string ", s)
	// 	return
	// }

	// fmt.Printf("Unknow Type %T\n", p)

	switch v := p.(type) {
	case string:
		fmt.Println("string", v)
	case int:
		fmt.Println("integer", v)
	default:
		fmt.Printf("%T\n", p)

	}

}

func TestEmptyInterfaceAssertion(t *testing.T) {

	DoSomething(10)
	DoSomething("hello")
	DoSomething(true)
}
