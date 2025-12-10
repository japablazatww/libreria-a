package liba

import (
	"errors"
	"fmt"
)

// GetUserBalance retrieves the balance for a user and account.
// It verifies the user ID and returns the balance.
func GetUserBalance(userID string, accountID string) (float64, error) {
	if userID == "" || accountID == "" {
		return 0, errors.New("userID and accountID are required")
	}
	// Simulation
	fmt.Printf("[Libreria-A] Getting balance for user %s, account %s\n", userID, accountID)
	return 1000.50, nil
}

// Transfer performs a money transfer between accounts.
// It takes source, destination, amount and checks for validity.
func Transfer(sourceAccount string, destAccount string, amount float64, currency string) (string, error) {
	if amount <= 0 {
		return "", errors.New("amount must be positive")
	}
	// Simulation
	fmt.Printf("[Libreria-A] Transferring %.2f %s from %s to %s\n", amount, currency, sourceAccount, destAccount)
	return "TX-123456789", nil
}

// GetSystemStatus checks the status of the system given an admin code.
// The code param is named simply "code" to test parameter mapping.
func GetSystemStatus(code string) (string, error) {
	if code != "ADMIN123" {
		return "Unauthorized", errors.New("invalid admin code")
	}
	// Simulation
	fmt.Printf("[Libreria-A] System Check by Admin %s\n", code)
	return "OPERATIONAL", nil
}
