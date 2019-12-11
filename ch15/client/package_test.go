package client

import (
	"testing"
	"ch15/serise"
)
func TestPackage(t *testing.T) {

	t.Log(serise.GetFibonacciSerise(5))
	t.Log(serise.Square(5))

}