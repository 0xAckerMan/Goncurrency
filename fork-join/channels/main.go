package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)
	<-done
	<-done
	<-done
	<-done
	fmt.Println("Finished waiting, exiting main")
	fmt.Println("Elapsed ", time.Since(now))
}

func task1(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1")
	done <- struct{}{}
}

func task2(done chan struct{}) {
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Task 2")
	done <- struct{}{}
}

func task3(done chan struct{}) {
	fmt.Println("Task 3")
	done <- struct{}{}
}

func task4(done chan struct{}) {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Task 4")
	done <- struct{}{}
}
