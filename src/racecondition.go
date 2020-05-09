package main

import (

	"fmt"
	//"runtime"
	"sync"
)

func main() {
	
	var wg sync.WaitGroup
	var m sync.Mutex 
	counter := 0
	gs := 100
	wg.Add(gs)
    for i:=0; i<gs; i++{
	    go func(){
		   m.Lock()
		   v := counter
		   //runtime.Gosched()
		   v++
		   counter = v
		   m.Unlock()
		   fmt.Println(counter)
		   //m.Unlock()
		   wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value",counter)
}
