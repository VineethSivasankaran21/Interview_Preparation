package patterns

import "fmt"

func Fan_Out_Pattern() {
	work := []int{33, 22, 44, 55, 66, 77, 88, 99, 111, 222, 333, 444}
	in := arrayGeneratorFanOut(work)

	out1 := workSeperator(in)
	out2 := workSeperator(in)
	out3 := workSeperator(in)
	out4 := workSeperator(in)
	out5 := workSeperator(in)
	out6 := workSeperator(in)
	out7 := workSeperator(in)

	for range in {
		select {
		case val := <-out1:
			fmt.Println("work has to be done by input1 : ", val)
		case val := <-out2:
			fmt.Println("work has to be done by input2 : ", val)
		case val := <-out3:
			fmt.Println("work has to be done by input3 : ", val)
		case val := <-out4:
			fmt.Println("work has to be done by input4 : ", val)
		case val := <-out2:
			fmt.Println("work has to be done by input2 : ", val)
		case val := <-out5:
			fmt.Println("work has to be done by input3 : ", val)
		case val := <-out5:
			fmt.Println("work has to be done by input4 : ", val)
		case val := <-out6:
			fmt.Println("work has to be done by input2 : ", val)
		case val := <-out7:
			fmt.Println("work has to be done by input3 : ", val)

		}

	}

}

func workSeperator(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for data := range in {
			out <- data
		}
	}()

	return out
}

func arrayGeneratorFanOut(work []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range work {
			ch <- v
		}
	}()
	return ch
}
