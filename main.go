package main

import "fmt"

func main() {
	event := []string{"teste-1", "teste-2", "teste-3", "teste-4"}
	event = event[1:]
	fmt.Println(event)
}
