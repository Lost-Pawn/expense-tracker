package main

import (
	"flag"
	"time"
	"fmt"
	"os"
	"expense-tracker/internal/storage"
	"expense-tracker/internal/task"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}
	
	loadExpenses, err := storage.LoadExpenses()
	if err != nil {
		fmt.Printf("Error loading expenses: %v\n", err)
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

		loadExpenses = append(loadExpenses, task.Expense{
			ID:          task.NextID(loadExpenses),
			Date:        time.Now().UTC(),
			Description: *description,
			Amount:      *amount,
		})

		err = storage.SaveExpenses(loadExpenses)
		if err != nil {
			fmt.Printf("Error saving expenses: %v\n", err)
			return
		} else {
			fmt.Printf("Adding expense: %s, Amount: %.2f\n", *description, *amount)

		}
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