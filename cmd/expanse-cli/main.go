package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Getting started with Expanse CLI")

	switch (os.Args[1]) {
	case "add":
		fmt.Println("add func")
	case "list":
		fmt.Println("list func")
	case "delete":
		fmt.Println("delete func")
	case "summary":
		fmt.Println("summary func")
	case "update":
		fmt.Println("update func")
	case "set-budget":
		fmt.Println("set-budget func")
	default:
		fmt.Println("Invalid command")
	}
}