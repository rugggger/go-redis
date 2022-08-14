package main

import "fmt"

func main() {

	fmt.Println("test")
	Put("test", "yaron")
	fmt.Println(Get("test"))
	Delete("test")
	fmt.Println(Get("test"))
	serve()

}
