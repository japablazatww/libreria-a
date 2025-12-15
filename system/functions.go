package system

import (
	"errors"
	"fmt"
)

// GetSystemStatus checks the status of the system given an admin code.
func GetSystemStatus(code string) (string, error) {
	if code != "ADMIN123" {
		return "Unauthorized", errors.New("invalid admin code")
	}
	// Simulation
	fmt.Printf("[Libreria-A] System Check by Admin %s\n", code)
	return "OPERATIONAL", nil
}
