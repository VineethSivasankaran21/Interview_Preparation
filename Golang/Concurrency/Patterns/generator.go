package patterns

import "fmt"

func GeneratorPattern() {
	c := generator()
	for val := range c {
		fmt.Println(val)
	}
}

func generator() <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 22; i < 27; i++ {
			out <- i
		}
	}()

	return out

}
