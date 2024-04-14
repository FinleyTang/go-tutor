package main

import "fmt"
import "math"

func test_value(){
	fmt.Println("============test_value================")
	fmt.Println("go"+"lang")

	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=",7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)


}

func test_variable(){
	fmt.Println("============test_variable================")
	var a = "initial"
	fmt.Println(a)
	
	var b, c int = 1, 2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	f := "short"
	fmt.Println(f)
}

const s string = "constant string"
func test_constant(){
	fmt.Println(s)
	const n = 500000
	const d = 3e20/n
	fmt.Println(d)
	fmt.Println(int(d))
	fmt.Println(math.Sin(n))
}


func test_loop()  {
	i := 1
	for i < 5{
		fmt.Println(i)
		i = i+1
	}
	count := 5
	for j := 0; j < count; j++ {
		fmt.Println(j)
	}

	arr := [2]int{1,2}
	for _, v := range arr {
		fmt.Println(v)
	}
}

func test_if()  {
	var a = 10
	if a > 11{
		fmt.Println("bigger")
	}else {
		fmt.Println("smaller")
	}
}

func main()  {
	fmt.Println("hello world")
	test_value()
	test_variable()
	test_constant()
	test_loop()
	test_if()
}
