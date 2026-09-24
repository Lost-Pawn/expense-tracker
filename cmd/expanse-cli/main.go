package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Getting started with Expanse CLI")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := addCmd.String("description", "", "Description of the expense")
		amount := addCmd.Float64("amount", 0, "Amount of the expense")
		addCmd.Parse(os.Args[2:])

		if *description == "" || *amount <= 0 {
			fmt.Println("Please provide a valid description and amount")
			return
		}

		fmt.Printf("Adding expense: %s, Amount: %f\n", *description, *amount)
		
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
	case "help":
		fmt.Println("help func")
	default:
		fmt.Println("Invalid command")
	}
}