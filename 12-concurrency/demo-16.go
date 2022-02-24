package main

import (
	"fmt"
	"time"
)

/* Generating even numbers for 5 seconds (using time.After() ) */

func main() {
	done := time.After(5 * time.Second)
	evenCh := genEvenNos(done)
	for no := range evenCh {
		fmt.Println(no)
	}
	fmt.Println("main completed")
}

func genEvenNos(done <-chan time.Time) <-chan int {
	var evenNoCh = make(chan int)
	go func() {
		no := 0
		for {
			select {
			case <-done:
				close(evenNoCh)
				return
			case evenNoCh <- no * 2:
				time.Sleep(500 * time.Millisecond)
				no++
			}
		}
	}()
	return evenNoCh
}
