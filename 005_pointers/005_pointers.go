package main

import "fmt"

func test_basic()  {
	var a int = 42
	var p *int
	p = &a
	fmt.Println("value of a:", a)
	fmt.Println("value of *p:", *p)
	fmt.Println("address of a:", &a)
	fmt.Println("value of p:", p)
}
//指针作为函数参数
func modifyValue(ptr *int)  {
	*ptr = 100
}

type Counter struct{
	count int
}

//指针接收者方法
func (c *Counter) Increment()  {
	c.count ++
}

func main()  {
	test_basic()
	value := 10
	modifyValue(&value)
	fmt.Println(value)

	var counter Counter
	fmt.Println(counter.count)
	counter.Increment()
	fmt.Println(counter.count)

}