package main

import "fmt"
import "reflect"

func test_arr_01()  {
	fmt.Println("==========test_arr========")
	var a[5] int
	fmt.Println("a: ", a)	

	b := [2]int{1,2}
	fmt.Println(b[1])

	var temparray [5]int
	for i := 0; i < 5; i++ {
		temparray[i] = i
	}
	for _, v := range temparray {
		fmt.Println(v)
	}
}

// interface{} 类型是一种特殊的类型，它可以容纳任意类型的值。
// 当你将一个值的类型声明为 interface{} 时，实际上是在说这个值可以是任何类型。
func check_type(v interface{})  {
		if reflect.TypeOf(v).Kind() == reflect.Slice{
			fmt.Println("you input is Slice")
		}else{
			fmt.Println("your input is not slice")
		}
}

func test_slice()  {
	fmt.Println("==========test_slices========")
	//You can create a slice using the make function or by slicing an existing array or another slice
	// This creates a slice of integers with a length of 5 and a capacity of 10. The capacity specifies the maximum size the slice can grow to before needing to allocate more memory.
	slice := make([]int, 5, 10)

	slice[0] = 1
	fmt.Println(slice[1], len(slice), cap(slice))

	// this is array
	arr := [5]int{1,2,3,4,5}
	check_type(arr)
	//赋值后的slice2 将是一个切片（slice），它指向了数组 arr 的相同底层数组
	slice2 := arr
	fmt.Println(slice2[1], len(slice2))
	// this is slice
	slice3 := []int{1,2,3,4,5}
	fmt.Println(len(slice3), cap(slice3))
	check_type(slice3)
	slice3 = append(slice3, 6)
	slice3 = append(slice3, 7, 8, 9)
	for _, v := range slice3 {
		fmt.Println(v)
	}



}

func main()  {
	fmt.Println("==========main========")
	test_arr_01()
	test_slice()
}