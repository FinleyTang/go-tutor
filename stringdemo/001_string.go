package main

import(
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("Contains: ", strings.Contains("test", "es"))
	fmt.Println("Count:    ", strings.Count("test", "t"))
	fmt.Println("HasPrefix:", strings.HasPrefix("test", "te"))
	fmt.Println("Split:    ", strings.Split("a-b-c", "-"))

}