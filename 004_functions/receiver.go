package main

import "fmt"

type Person struct {
	age int
}

func (p Person) howOld()  {
	p.age = p.age +1
	fmt.Println("how old inner:",p.age)
}

func (p *Person) growUp() {
	p.age += 1
}

func main() {
	// qcrao 是值类型
	qcrao := Person{age: 18}

	// 值类型 调用接收者也是值类型的方法
	qcrao.howOld()
	fmt.Print(qcrao.age)

}