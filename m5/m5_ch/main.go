package main

import (
	"fmt"
	"time"
)

func main() {
	semaphore := make(chan int, 3)
	for i := 0; i < 10; i++ {
		semaphore <- 1
		go func(){
			defer func(){
				msg := <-semaphore
				fmt.Println(msg)
			}()
			time.Sleep(time.Millisecond * 1000)
		}()
	}
	for len(semaphore) > 0 {
		time.Sleep(time.Millisecond * 10)
	}
}


/*func main() {
	r := make(chan float64)
	go func() {	
		val, ok := <-r
	fmt.Println(val, ok)
	}()
	r <- 1.2
	
	fmt.Println("success")
}
*/
/*func main() {
	ch1 := make(chan int)
	fmt.Println(len(ch1), cap(ch1))
	ch2 := make(chan int, 3)
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 4
	fmt.Println(len(ch2), cap(ch2))

	close(ch1)
	result, ok := <- ch1
	fmt.Println(result, ok)
	result, ok = <- ch2
	fmt.Println(result, ok)
}*/