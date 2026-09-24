package task

import (
	"time"
)

type Expense struct {
    ID          int     	`json:"id"`
    Date        time.Time 	`json:"date"`
    Description string  	`json:"description"`
    Amount      float64 	`json:"amount"`
}

func NextID(expenses []Expense) int {
    maxID := 0
    for _, expense := range expenses {
        if expense.ID > maxID {
            maxID = expense.ID
        }
    }
    return maxID + 1
}