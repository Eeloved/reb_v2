package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)
	go spinner(100 * time.Millisecond)
	//go work()
	//wg.Wait()

go peral(&wg)

wg.Wait()
elapsed := time.Since(start)
fmt.Printf("%s\n", elapsed)
}
	//n := 45
func peral(wg *sync.WaitGroup) {
	
	x := [...]int{44, 45}
	for i := 0; i < len(x); i++ {
		
		go func(n int){
		defer wg.Done()
		fibN := fib(x[i])
		fmt.Printf("\rFibonacci(%d) = %d\n", x[i], fibN)
	}(i)
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/`  {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2{
		return x
	}
	return fib(x-1) + fib(x-2)
}
