package main

import "sync"

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	go func() {
		defer close(ch)
		for item := range ch {
			println(item)
			wg.Done()
		}
	}()
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		ch <- i
	}
	wg.Wait()

}
