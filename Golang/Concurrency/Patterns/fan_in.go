package patterns

import (
	"fmt"
	"sync"
)

func FanInPattern() {
	in1 := arrayGenerator([]int{21, 22, 23, 24, 25})
	in2 := arrayGenerator([]int{31, 32, 33, 34, 35})

	out := outPut(in1, in2)

	for val := range out {
		fmt.Println("Value : --> ", val)
	}
}

func outPut(input ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(len(input))

	for _, val := range input {
		go func(in <-chan int) {
			for {
				v, ok := <-in
				if !ok {
					wg.Done()
					break
				}
				out <- v
			}
		}(val)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func arrayGenerator(val []int) <-chan int {
	in := make(chan int)

	go func() {
		defer close(in)
		for _, v := range val {
			in <- v
		}
	}()

	return in
}
