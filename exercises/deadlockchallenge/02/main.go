package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	for x := range c {
		fmt.Println(x)
	}

}

//Original
/*
C := make(chan int)
go func(){
  for i := 0; i <10; i++{
    c <- i
}
}()
fmt.Println(<-c)
*/
