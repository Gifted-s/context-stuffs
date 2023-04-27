package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	countChan := make(chan int)
	ctx, cancelChannel := context.WithCancel(ctx)
	go doAnotherThing(ctx, countChan)

	for num := 0; num <= 3; num++ {
		countChan <- num
	}

	cancelChannel()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("doSomething: finished")

}

func doAnotherThing(ctx context.Context, countChan <-chan int) {
	for {
		select {
		case <-ctx.Done():

			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnotherThing err: %s\n", err)
				fmt.Println("doAnotherThing Finished")
				return
			}

		case num := <-countChan:
			fmt.Printf("Count: %d\n", num)
		}
	}

}

func main() {
	ctx := context.Background()

	doSomething(ctx)

}
