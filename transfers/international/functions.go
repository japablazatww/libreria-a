package international

import (
	"errors"
	"fmt"
)

// InternationalTransfer performs a cross-border transfer.
func InternationalTransfer(sourceAccount string, destIban string, amount float64, swiftCode string) (string, error) {
	if amount <= 0 {
		return "", errors.New("amount must be positive")
	}
	// Simulation
	fmt.Printf("[Libreria-A/International] Sending %.2f to %s (SWIFT: %s)\n", amount, destIban, swiftCode)
	return "INT-TX-999999", nil
}
