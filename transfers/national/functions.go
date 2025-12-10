package national

import (
	"errors"
	"fmt"
)

// GetUserBalance retrieves the balance for a user and account.
func GetUserBalance(userID string, accountID string) (float64, error) {
	if userID == "" || accountID == "" {
		return 0, errors.New("userID and accountID are required")
	}
	// Simulation
	fmt.Printf("[Libreria-A/National] Getting balance for user %s, account %s\n", userID, accountID)
	return 1000.50, nil
}

// Transfer performs a local money transfer.
func Transfer(sourceAccount string, destAccount string, amount float64, currency string) (string, error) {
	if amount <= 0 {
		return "", errors.New("amount must be positive")
	}
	// Simulation
	fmt.Printf("[Libreria-A/National] Transferring %.2f %s from %s to %s\n", amount, currency, sourceAccount, destAccount)
	return "TX-123456789", nil
}
