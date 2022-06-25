package main

import "fmt"

func main() {

	ch := make(chan int)
	ch1 := make(chan int)
	out := make(chan int)

	// Channel 1
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i + 1
			fmt.Println("produced ch one", i+1)
		}
		close(ch)
	}()

	// Channel 2
	go func() {
		for i := 11; i < 21; i++ {
			ch1 <- i + 1
			fmt.Println("produced ch two", i+1)
		}
		close(ch1)
	}()

	go Produce(ch, ch1, out)

	for v := range out {
		fmt.Println("consumed ch three", v)
	}

}

//Channel 3
func Produce(ch, ch1, out chan int) {
	for v := range ch {
		out <- v
	}
	for v := range ch1 {
		out <- v
	}
	close(out)
}
