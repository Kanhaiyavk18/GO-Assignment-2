package main

import "fmt"

func main() {

	channel := make(chan int)
	channel1 := make(chan int)
	out := make(chan int)

	// Ch 1
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i + 1
			fmt.Println("produced channel one", i+1)
		}
		close(channel)
	}()

	// Ch 2
	go func() {
		for i := 11; i < 21; i++ {
			channel1 <- i + 1
			fmt.Println("produced channel two", i+1)
		}
		close(channel1)
	}()

	go Produce(channel, channel1, out)

	for v := range out {
		fmt.Println("consumed channel three", v)
	}

}

//Ch 3
func Produce(channel, channel1, out chan int) {
	for v := range channel {
		out <- v
	}
	for v := range channel1 {
		out <- v
	}
	close(out)
}
