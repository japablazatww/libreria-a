package national

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
