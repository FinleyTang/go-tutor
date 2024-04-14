package main

import(
	"fmt"
	"os"
	"bufio"
)

func check_err(e error)  {
	if e !=nil {
		panic(e)
	}
}

func writeFile(filename string)  {
	file, err := os.Create(filename)
	check_err(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("hello world")
	check_err(err)
	err = writer.Flush()
	check_err(err)

	fmt.Println("file write successfully")
}

func main()  {
	writeFile("hello.txt")	
}