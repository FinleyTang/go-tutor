package main

import "fmt"

type address struct{
	country string 
	city string
}

type person struct{
	name string 
	age int
	addr address
}

func (p *person)setname(name string)  {
	p.name = name
}


func main()  {
	p := person{name:"ftang", age:30}
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)

	//注意这里的写法
	p2 := person{name:"alice", age:21, addr:address{country:"US", city:"new york"}}
	fmt.Println(p2.name, p2.age, p2.addr.country, p2.addr.city)

	p2.setname("bob")
	fmt.Println(p2.name)

}