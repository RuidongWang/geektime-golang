// channel 关闭 demo
package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			if data, ok := <-ch; ok {
				fmt.Println("data is : ", data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

/**
 * 向关闭的 channel 发送数据，会导致 panic
 * v, ok <- ch; ok 为 bool 值，true 表示正常接受，false 表示通道关闭
 * 所有的 channel 接收者都会在 channel 关闭时，立即从阻塞等待中返回且上述 ok 值为 false。这个广播机制常被利用，进行向多个订阅者同时发送信号。如：退出信号
 */

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
