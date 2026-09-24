package task

import (
	"time"
    "fmt"
)

type Expense struct {
    ID          int     	`json:"id"`
    Date        time.Time 	`json:"date"`
    Description string  	`json:"description"`
    Amount      float64 	`json:"amount"`
}

type Budget struct {
    Month   int     `json:"month"`
    Amount  float64 `json:"amount"`
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

func CheckBudget(expenses []Expense, budgets []Budget, month int) {
    var total float64
    for _, expense := range expenses {
        if int(expense.Date.Month()) == month {
            total += expense.Amount
        }
    }

    for _, budget := range budgets {
        if budget.Month == month {
            if total > budget.Amount {
                fmt.Printf("Warning: You have exceeded your budget for month %d! (%.2f / %.2f)\n", month, total, budget.Amount)
            }
            return
        }
    }
}