package lunchmoney

import "fmt"

// RecurringExpense ...
type RecurringExpense struct {
	ID             string `json:"id"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	Cadence        string `json:"cadence"`
	Payee          string `json:"payee"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	Description    string `json:"description"`
	BillingDate    string `json:"billing_date"`
	Type           string `json:"type"`
	OriginalName   string `json:"original_name"`
	Source         string `json:"source"`
	TransactionID  string `json:"transaction_id"`
	CategoryID     int64  `json:"category_id"`
	AssetID        int64  `json:"asset_id"`
	PlaidAccountID int64  `json:"plaid_account_id"`
}

// GetRecurringExpensesOptions ...
type GetRecurringExpensesOptions struct {
	StartDate       string
	DebitAsNegative bool
}

// GetRecurringExpensesResponse ...
type GetRecurringExpensesResponse struct {
	RecurringExpenses []RecurringExpense `json:"recurring_expenses"`
}

func getRecurringExpensesQuery(opts *GetRecurringExpensesOptions) string {
	// To Implement
	return ""
}

// GetRecurringExpenses ...
func (client *Client) GetRecurringExpenses(opts *GetRecurringExpensesOptions) (*GetRecurringExpensesResponse, error) {

	resp := GetRecurringExpensesResponse{}
	queries := ""

	if opts != nil {
		queries = getRecurringExpensesQuery(opts)
	}

	err := client.Call("GET", fmt.Sprintf("recurring_expenses%s", queries), nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
