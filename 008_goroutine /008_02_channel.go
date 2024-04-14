package main

import "fmt"


func sum(arr[]int, c chan int )  {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func counter(c chan int)  {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}


func main()  {
	s :=[]int{1,2,3,4,5,6}	

	c := make(chan int)
	go sum(s[:3], c)
	go sum(s[3:], c)
	x, y := <-c, <-c
	fmt.Println(x)	
	fmt.Println(y)	

	ch := make(chan int)
	go counter(ch)

	for  v := range ch {
		fmt.Println(v)
	}

}