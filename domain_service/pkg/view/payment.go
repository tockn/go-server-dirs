package view

type PayRequest struct {
	PayerID int64 `json:"payerID"`
	PayeeID int64 `json:"payeeID"`
	Amount  int64 `json:"amount"`
}

type PayResponse struct {
	Status        string `json:"status"`
	FailedMessage string `json:"failed_message,omitempty"`
}
