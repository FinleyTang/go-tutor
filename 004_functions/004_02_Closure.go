package main

import "fmt"

//当我们说一个函数是一个闭包时，意味着这个函数可以访问其外部作用域中的变量。这意味着在函数内部定义的变量可以被该函数引用和操作，
//即使这些变量在函数外部定义，甚至在函数调用完成后仍然可以被访问和修改。
func outFunction() func(){
	varA := 10
	result := 0
	return func ()  {
		result = result + varA
		fmt.Println(result)
	}

}

func main()  {
	closure_func := outFunction()
	closure_func()
	closure_func()
	closure_func()

	//在这个例子中，outerFunction 返回了一个闭包，闭包内部可以访问外部作用域中的 outerVariable 变量。即使 outerFunction 执行完毕后，闭包仍然可以访问和操作 outerVariable 变量。
	//这是因为闭包不仅包含了函数体，还包含了其外部作用域中的变量的引用。
}