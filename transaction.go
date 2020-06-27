package lunchmoney

import "errors"

// Transaction ...
type Transaction struct {
	ID             int64  `json:"id"`
	Date           string `json:"date"`
	Payee          string `json:"payee"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	Notes          string `json:"notes"`
	CategoryID     int64  `json:"category_id"`
	AssetID        int64  `json:"asset_id"`
	PlaidAccountID int64  `json:"plaid_account_id"`
	Status         string `json:"status"`
	ParentID       int64  `json:"parent_id"`
	IsGroup        bool   `json:"is_group"`
	GroupID        int64  `json:"group_id"`
	Tags           []Tag  `json:"tags"`
	ExternalID     string `json:"external_id"`
}

// Tag ...
type Tag struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Options ...
type Options struct {
	TagID           int64
	RecurringID     int64
	PlaidAccountID  int64
	CategoryID      int64
	AssetID         int64
	Offset          int64
	Limit           int64
	DebitAsNegative bool
}

// GetTransactionsResponse ...
type getTransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

// GetTransactions ...
func (client *Client) GetTransactions(startDate, endDate string, opts *Options) (transactions *[]Transaction, err error) {

	if startDate == "" || endDate == "" {
		return nil, errors.New("Start and End Dates must be specified")
	}

	resp := getTransactionsResponse{}

	err = client.Call("GET", "transactions", nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Transactions, nil
}
