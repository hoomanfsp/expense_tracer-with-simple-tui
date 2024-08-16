package proceed

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// Expense struct should match the struct defined in the 'database' package
type Expense struct {
	ID          uint    `gorm:"primary_key"`
	Amount      float64 // Amount should be float64
	Description string  // Description should be a string
}

// Add inserts a new expense into the database
func Add(amount, description string, db *gorm.DB) {
	value := Expense{
		Amount:      parseAmount(amount), // Convert the amount to a float64
		Description: description,
	}
	db.Create(&value)
}

// Delete removes an expense by its ID from the database
func Delete(id string, db *gorm.DB) {
	db.Delete(&Expense{}, id)
}

// List retrieves all expenses from the database and returns them as a formatted string,
// including the sum of all amounts
func List(db *gorm.DB) string {
	var expenses []Expense
	db.Find(&expenses)

	var sb strings.Builder
	var totalAmount float64

	for _, expense := range expenses {
		sb.WriteString(fmt.Sprintf("ID: %d | Amount: %.2f | Description: %s\n", expense.ID, expense.Amount, expense.Description))
		totalAmount += expense.Amount
	}

	// Add the sum of all amounts at the end
	sb.WriteString(fmt.Sprintf("\nTotal Amount: %.2f", totalAmount))

	return sb.String()
}

// parseAmount converts the amount string to float64
func parseAmount(amount string) float64 {
	var result float64
	fmt.Sscanf(amount, "%f", &result)
	return result
}
