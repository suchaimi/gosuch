package main

import "fmt"

func main() {
	var command = "restart"

	if command == "restart" {
		fmt.Println("Hell just restart the computer")
	} else if command == "shutdown" {
		fmt.Println("Hell just power off the compputer")
	} else {
		fmt.Println("Hell dont know the command")
	}
}
