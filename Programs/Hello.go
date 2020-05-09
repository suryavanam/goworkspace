package main

import (
	"fmt"
	"runtime"
	"sync"
)

//func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Number of cpus start", runtime.NumCPU())
	fmt.Println("num of Go routines start", runtime.NumGoroutine())
	go func() {
		fmt.Println("This is anonymus first")
		wg.Done()

	}()

	go func() {
		fmt.Println("This is anonymus second")
		wg.Done()
	}()
	fmt.Println("Number of cpus mid", runtime.NumCPU())
	fmt.Println("num of Go routines mid", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Number of cpus end", runtime.NumCPU())
	fmt.Println("num of Go routines end", runtime.NumGoroutine())
	fmt.Println("Program exits")

}
