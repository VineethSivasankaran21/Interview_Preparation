package patterns

import "fmt"

func PipelinePattern() {
	in := arrayGenerator([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	out := seperateOdd(in)
	out = square(out)
	out = Half(out)

	for val := range out {
		fmt.Println(val)
	}
}

func seperateOdd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in {
			if val%2 != 0 {
				out <- val
			}
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in {
			out <- val * val
		}
	}()
	return out
}

func Half(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in {
			out <- val / 2
		}
	}()
	return out
}
