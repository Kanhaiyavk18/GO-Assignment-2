package main

import "fmt"

func main() {

	var size int
	fmt.Print("Enter size of array : ")
	fmt.Scanf("%d",&size)

//consumeChannel fn to consume ch
	
     channel := consumeChannel() 

// for Producing ch
		for i := 0; i <= size; i++{
		channel <- i+1
		fmt.Println("Produced",i+1)
	}
}
func consumeChannel() chan <- int {

	channel := make(chan int)

	go func() {
		for k:= range channel {
			fmt.Println("Consumed",k)
		
		}
	}()

	return channel
}