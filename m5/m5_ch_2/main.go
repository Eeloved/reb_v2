package main

import(
	"context"
	"fmt"
	//"time"
)


func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func(){
			for{
				select {
					case <- ctx.Done():
						return
					case dst <- n:
						n++
					}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5{
			break
		}
	}
}

/*
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 300)
	tick := time.NewTicker(time.Millisecond * 50)

	for {
		select {
		case t := <-tick.C:
			fmt.Println(t)
			cancel()
		case <-ctx.Done():
			fmt.Println("context deadline exceeded")
			return
		default:
			fmt.Println("waiting")
			time.Sleep(time.Millisecond * 20)

		}
	}
}*/

/*
func main() {
	semaphore := make(chan int, 3)
	done := make(chan struct{})
	i := 0

	go func() {
		for  ; ; i++ {
			semaphore <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		time.Sleep(time.Millisecond * 1000)
		done <- struct{}{}
	}()
	msg := 0
L:
	for {
		select {
		case msg = <-semaphore:
			fmt.Println(msg)
		case <- done:
			fmt.Println("done")
			break L
		default:
			if msg >= 20 {
				break L
			}
			fmt.Println("waiting")
			time.Sleep(time.Millisecond * 200)
		}
		
	}
	fmt.Println("success")
}*/