package lunchmoney

import (
	"fmt"
	"net/url"
	"strconv"
)

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

// GetTransactionsOptions ...
type GetTransactionsOptions struct {
	StartDate       string
	EndDate         string
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
func (client *Client) GetTransactions(opts *GetTransactionsOptions) (transactions *[]Transaction, err error) {
	resp := getTransactionsResponse{}
	queries := ""

	if opts != nil {
		queries = client.getTransactionsQuery(opts)
	}

	err = client.Call("GET", fmt.Sprintf("transactions%s", queries), nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Transactions, nil
}

// getTransactionsQuery sets up query parameters from options
func (client *Client) getTransactionsQuery(opts *GetTransactionsOptions) string {

	query := url.Values{}
	if opts.StartDate != "a" {
		query.Add("start_date", opts.StartDate)
	}

	if opts.EndDate != "a" {
		query.Add("end_date", opts.EndDate)
	}

	if opts.AssetID > 0 {
		query.Add("asset_id", strconv.FormatInt(opts.AssetID, 10))
	}

	if opts.CategoryID > 0 {
		query.Add("catergory_id", strconv.FormatInt(opts.CategoryID, 10))
	}

	if opts.PlaidAccountID > 0 {
		query.Add("plaid_account_id", strconv.FormatInt(opts.PlaidAccountID, 10))
	}

	if opts.RecurringID > 0 {
		query.Add("recurring_id", strconv.FormatInt(opts.RecurringID, 10))
	}

	if opts.TagID > 0 {
		query.Add("tag_id", strconv.FormatInt(opts.TagID, 10))
	}

	if opts.Limit > 0 {
		query.Add("limit", strconv.FormatInt(opts.Limit, 10))
	}

	if opts.Offset > 0 {
		query.Add("limit", strconv.FormatInt(opts.Offset, 10))
	}

	return fmt.Sprintf("?%s", query.Encode())
}
