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

type TransferRequest struct {
	SourceAccount string  `json:"source_account"`
	DestAccount   string  `json:"dest_account"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
}

type TransferResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
}

// ComplexTransfer performs a transfer using a struct input/output.
func ComplexTransfer(req TransferRequest) (TransferResponse, error) {
	if req.Amount <= 0 {
		return TransferResponse{Status: "FAILED"}, errors.New("amount must be positive")
	}
	fmt.Printf("[Libreria-A/National] ComplexTransfer: %v\n", req)
	return TransferResponse{
		TransactionID: "CTX-99999",
		Status:        "SUCCESS",
	}, nil
}
