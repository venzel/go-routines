package channels

import (
	"fmt"
	"sync"
)

func publish(ch chan int, count int) {
	for i := 0; i < count; i++ {
		ch <- i
	}
	close(ch)
}

func read(ch chan int, count int, wg *sync.WaitGroup) {
	wg.Add(count)
	for i := range ch {
		fmt.Println(i)
		wg.Done()
	}
}

func Execute() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	count := 100

	go publish(ch, count)
	read(ch, count, &wg)

	wg.Wait()
}
