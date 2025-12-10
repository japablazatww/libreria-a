package national

type NationalTransfers interface {
	GetUserBalance(userID string, accountID string) (float64, error)
	Transfer(sourceAccount string, destAccount string, amount float64, currency string) (string, error)
}
