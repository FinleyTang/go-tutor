package main

import "fmt"

//// 定义一个接受变长参数的函数，参数类型为 int
func test_variadic(nums ...int) int {
	var result int
	for _, v := range nums {
		result += v		
	}
	return result
}


func main()  {
	arr := []int{1,2,3,4,5}
	//// 可以传入一个切片作为参数
	result := test_variadic(arr...)
	fmt.Println(result)

	arr2 := [5]int{1, 2, 3, 4, 5}
    fmt.Println(test_variadic(arr2[0], arr2[1], arr2[2], arr2[3], arr2[4])) // 正确
	// fmt.Println(test_variadic(arr2)) //会产生错误
	// 转化为slice则可以
	fmt.Println(test_variadic(arr2[:]...))

}