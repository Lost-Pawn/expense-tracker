package storage

import (
	"encoding/json"
	"errors"
	"io/fs"
	"expense-tracker/internal/task"
	"fmt"
	"os"
	"path/filepath"
)

func LoadExpenses() ([]task.Expense, error) {
	expensesPath := filepath.Join(".", "expenses.json")
	file, err := os.Open(expensesPath)
	if errors.Is(err, fs.ErrNotExist) {
		return []task.Expense{}, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to check expenses file: %w", err)
	}
	defer file.Close()
	
	var expenses []task.Expense
	err = json.NewDecoder(file).Decode(&expenses)
	if err != nil {
		return nil, fmt.Errorf("failed to decode expenses: %w", err)
	}

	return expenses, nil
}

func SaveExpenses(expenses []task.Expense) error {
	expensesPath := filepath.Join(".", "expenses.json")
	file, err := os.Create(expensesPath)
	if err != nil {
		return fmt.Errorf("failed to create expenses file: %w", err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(expenses)
	if err != nil {
		return fmt.Errorf("failed to encode expenses: %w", err)
	}

	return nil
}