package xfers

import (
	"context"
	"fmt"
	"net/http"
)

// Balance is account balance model.
type Balance struct {
	TotalBalance     string `json:"totalBalance"`
	AvailableBalance string `json:"availableBalance"`
	PendingBalance   string `json:"pendingBalance"`
}

// GetBalance to get account balance.
func (c *Client) GetBalance() (*Balance, int, error) {
	return c.GetBalanceWithContext(context.Background())
}

// GetBalanceWithContext to get account balance with context.
func (c *Client) GetBalanceWithContext(ctx context.Context) (*Balance, int, error) {
	var response balance
	code, err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/overviews/balance_overview", c.baseURL),
		c.apiKey,
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, code, err
	}
	return &response.Data.Attributes, code, nil
}

// Bank is disbursement bank model.
type Bank struct {
	Name      string   `json:"name"`
	ShortCode BankCode `json:"shortCode"`
}

// GetBanks to get disbursement bank list.
func (c *Client) GetBanks() ([]Bank, int, error) {
	return c.GetBanksWithContext(context.Background())
}

// GetBanksWithContext to get disbursement bank list with context.
func (c *Client) GetBanksWithContext(ctx context.Context) ([]Bank, int, error) {
	var response bank
	code, err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/banks", c.baseURL),
		c.apiKey,
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, code, err
	}
	return response.toBanks(), code, nil
}

// BankAccount is bank account model.
type BankAccount struct {
	AccountName   string   `json:"accountName"`
	AccountNo     string   `json:"accountNo"`
	BankShortCode BankCode `json:"bankShortCode"`
}

// ValidateBankAccount to validate bank account.
func (c *Client) ValidateBankAccount(request ValidateBankAccountRequest) (*BankAccount, int, error) {
	return c.ValidateBankAccountWithContext(context.Background(), request)
}

// ValidateBankAccountWithContext to validate bank account with context.
func (c *Client) ValidateBankAccountWithContext(ctx context.Context, request ValidateBankAccountRequest) (*BankAccount, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response bankAccount
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/validation_services/bank_account_validation", c.baseURL),
		c.apiKey,
		c.secretKey,
		nil,
		request.wrap(),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return &response.Data.Attributes, code, nil
}
