package main

import "sync"

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go feeder(ch, &wg)
	go catcher(ch)
	close(ch)
	wg.Wait()
}

func catcher(ch chan int) {
	for {
		foo, ok := <-ch
		if !ok {
			println("done")
			return
		}
		println(foo)
	}
	wg.Done()
}

func feeder(ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; i <= 5; i++ {

		ch <- i
	}
	wg.Done()
}
