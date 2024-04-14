package main

import(
	"fmt"
	"os"
	"bufio"
)

func check_err(e error)  {
	if e !=nil{
		panic(e)
	}
}

func readFile(filename string){
	data, err := os.ReadFile(filename)
	check_err(err)
	fmt.Println(string(data))
}

func readFileLineByLine(filename string)  {
	f, err := os.Open(filename)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}	
	err = scanner.Err()
	check_err(err)
}


func readFileLine(filename string)  {
	f, err := os.Open(filename)
	check_err(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	err = scanner.Err()
	check_err(err)

}


func main()  {
	readFile("test.txt")
	// readFileLineByLine("test.txt")
	readFileLine("test.txt")
}