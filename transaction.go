package lunchmoney

import (
	"encoding/json"
	"errors"
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
type GetTransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

// getTransactionsQuery sets up query parameters from options
func getTransactionsQuery(opts *GetTransactionsOptions) string {

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

// GetTransactions ...
func (client *Client) GetTransactions(opts *GetTransactionsOptions) (*GetTransactionsResponse, error) {

	resp := GetTransactionsResponse{}
	queries := ""

	if opts != nil {
		queries = getTransactionsQuery(opts)
	}

	err := client.Call("GET", fmt.Sprintf("transactions%s", queries), nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetTransactionByID ...
func (client *Client) GetTransactionByID(transactionID int64) (*Transaction, error) {
	resp := Transaction{}
	endpoint := fmt.Sprintf("transactions/%d", transactionID)

	err := client.Call("GET", endpoint, nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// InsertTransactionsOptions ...
type InsertTransactionsOptions struct {
	ApplyRules      bool
	CheckRecurring  bool
	DebitAsNegative bool
}

// InsertTransactionsRequest ...
type InsertTransactionsRequest struct {
	Transactions []Transaction `json:"transactions"`
}

// InsertTransactionsResponse ...
type InsertTransactionsResponse struct {
	IDs    []int64  `json:"ids"`
	Errors []string `json:"error"`
}

func insertTransactionsQuery(opts *InsertTransactionsOptions) string {
	// To Implement
	return ""
}

// InsertTransactions ...
func (client *Client) InsertTransactions(transactions []Transaction, opts *InsertTransactionsOptions) (*InsertTransactionsResponse, error) {

	resp := InsertTransactionsResponse{}
	toReq := InsertTransactionsRequest{
		Transactions: transactions,
	}

	body, err := json.Marshal(toReq)
	if err != nil {
		return nil, err
	}

	err = client.Call("POST", "transactions", body, &resp)
	if err != nil {
		return nil, err
	}

	// Workaround -> Currently, the API returns 200 even if there are some types of
	// errors in the response.
	if len(resp.Errors) > 0 {
		return &resp, errors.New("Request errors, check response body")
	}

	return &resp, nil
}
