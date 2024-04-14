package main

import "fmt"

func test_map()  {
	map1 := make(map[string]string)
	map1["apple"] = "red"

	fmt.Println(map1["apple"])
	fmt.Println(map1)

	map2 := map[string]int{
		"a":1, 
		"b":2,
	}

	fmt.Println(map2)
	// check if a key exsit in map2
	_, prs := map2["c"]
    fmt.Println("prs:", prs) //print:  false
	for k, v := range map2 {
		fmt.Println(k,v)
	}
	//在 Go 语言中，map 字面量（以及其他复合数据类型如数组、切片等的字面量）的语法允许在最后一个元素后面添加一个额外的逗号。
	// 如果是花括号换行必须在最后加逗号
	// 注：
	person := map[string]int{
        "Alice": 30,
        "Bob":   25,
        "Cathy": 35, //花括号换行此处必须加逗号
    }
	for name, age := range person {
        fmt.Printf("%s is %d years old\n", name, age)
    }

}


func main()  {
	test_map()
}