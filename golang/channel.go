package main

import (
	"fmt"
	"sync"
)

/*
go 语言中, 当缓冲区 channel 关闭后, 若数据未全部读取完成,
读取操作会返回该 channel 元素类型的零值, 并且 ok 为 false.

优雅关闭 channel 是指通过 close 函数来通知所有接收方 channel
已经没有数据可以接受, 并且接收方在接收到 channel 关闭信号后,
不会再有新的数据发送进来. 避免接收方 goroutine 陷入永久阻塞状态.
*/

func channelClose() {
	numReaders := 2
	dataChan := make(chan int, 10) // 创建一个带缓冲的channel
	var wg sync.WaitGroup
	wg.Add(numReaders + 1) // 增加两个读取协程和一个写入协程的计数

	// 写入协程
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			dataChan <- i // 写入数据
		}
		close(dataChan) // 关闭channel
		fmt.Println("Writer goroutine finished and closed the channel")
	}()

	// 读取协程1
	go func() {
		defer wg.Done()
		for {
			data, ok := <-dataChan
			if !ok {
				fmt.Println("Reader 1: Channel closed")
				break
			}
			fmt.Printf("Reader 1: Received %d\n", data)
		}
	}()

	// 读取协程2
	go func() {
		defer wg.Done()
		for {
			data, ok := <-dataChan
			if !ok {
				fmt.Println("Reader 2: Channel closed")
				break
			}
			fmt.Printf("Reader 2: Received %d\n", data)
		}
	}()

	wg.Wait() // 等待所有协程结束
	fmt.Println("All goroutines finished")
}

func main() {

	channelClose()

}
