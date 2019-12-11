package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singletonInstance *Singleton

var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Println(unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
