package main

import (
	"expense-tracker/internal/storage"
	"expense-tracker/internal/task"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
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
		if len(loadExpenses) == 0 {
			fmt.Println("No expenses found")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "ID\tDate\tDescription\tAmount")

		for _, expense := range loadExpenses {
			fmt.Fprintf(w, "%d\t%s\t%s\t%.2f\n", expense.ID, expense.Date.Format("2006-01-02"), expense.Description, expense.Amount)
		}
		w.Flush()
	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "ID of the expense to delete")
		deleteCmd.Parse(os.Args[2:])

		if *id <= 0 {
			fmt.Println("Please provide a valid ID")
			return
		}

		for i, expenses := range loadExpenses{
			if expenses.ID == *id {
				loadExpenses = append(loadExpenses[:i], loadExpenses[i+1:]...)
				err = storage.SaveExpenses(loadExpenses)
				if err != nil {
					fmt.Printf("Error saving expenses: %v\n", err)
					return
				} else {
					fmt.Printf("Deleted expense with ID: %d\n", *id)
					return
				}
			}
		}
		fmt.Printf("No expense found with ID:%d\n", *id)
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