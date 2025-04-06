package events

type AccountAddMoneyEvent struct {
	UserID    string  `json:"userID"`
	Username  string  `json:"username"`
	AccountID string  `json:"accountID"`
	Ammount   float64 `json:"ammount"`
	Currency  string  `json:"currency"`
}

type AccountWithldrawEvent struct {
	UserID    string  `json:"userID"`
	Username  string  `json:"username"`
	AccountID string  `json:"accountID"`
	Ammount   float64 `json:"ammount"`
	Currency  string  `json:"currency"`
}
