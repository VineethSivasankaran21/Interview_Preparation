package basics

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mutex sync.Mutex
}

func ExampleMutex() {
	var wg sync.WaitGroup
	c := Counter{}
	wg.Add(4)
	c.Add(100, &wg)
	c.Add(-50, &wg)
	c.Add(-50, &wg)
	c.Add(3, &wg)
	wg.Wait()

}

func (c *Counter) Add(n int, wg *sync.WaitGroup) {
	c.mutex.Lock()
	defer wg.Done()
	Print(n, c.count)
	c.count += n
	c.mutex.Unlock()
}

func Print(n, c int) {
	switch {
	case n > 0:
		fmt.Printf("current value of counter : %d , Going to add : %d \n", c, n)
	case n < 0:
		fmt.Printf("current value of counter : %d , Going to subtract : %d \n", c, n)
	}

}
