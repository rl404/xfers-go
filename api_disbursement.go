package xfers

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Disbursement is disbursement model.
type Disbursement struct {
	ID                          string
	ReferenceID                 string
	Type                        DisbursementType
	Amount                      float64
	Fees                        float64
	Status                      Status
	BankAccountNo               string
	BankShortCode               BankCode
	BankName                    string
	BankAccountHolderName       string
	ServerBankAccountHolderName string
	Description                 string
	FailureReason               string
	CreatedAt                   time.Time
}

// CreateDisbursement to create new disbursement.
func (c *Client) CreateDisbursement(request CreateDisbursementRequest) (*Disbursement, int, error) {
	return c.CreateDisbursementWithContext(context.Background(), request)
}

// CreateDisbursementWithContext to create new disbursement with context.
func (c *Client) CreateDisbursementWithContext(ctx context.Context, request CreateDisbursementRequest) (*Disbursement, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response disbursement
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/disbursements", c.baseURL),
		c.apiKey,
		c.secretKey,
		nil,
		request.wrap(),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return response.toDisbursement(), code, nil
}

// GetDisbursement to get disbursement.
func (c *Client) GetDisbursement(id string) (*Disbursement, int, error) {
	return c.GetDisbursementWithContext(context.Background(), id)
}

// GetDisbursementWithContext to get disbursement with context.
func (c *Client) GetDisbursementWithContext(ctx context.Context, id string) (*Disbursement, int, error) {
	if id == "" {
		return nil, http.StatusBadRequest, errRequiredField("id")
	}

	var response disbursement
	code, err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/disbursements/%s", c.baseURL, id),
		c.apiKey,
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, code, err
	}
	return response.toDisbursement(), code, nil
}

// GetDisbursements to get disbursement list.
func (c *Client) GetDisbursements(request Pagination) ([]Disbursement, int, error) {
	return c.GetDisbursementsWithContext(context.Background(), request)
}

// GetDisbursementsWithContext to get disbursement list with context.
func (c *Client) GetDisbursementsWithContext(ctx context.Context, request Pagination) ([]Disbursement, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response disbursements
	code, err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/disbursements?%s", c.baseURL, request.encode()),
		c.apiKey,
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return response.toDisbursements(), code, nil
}

// DisbursementAction is response model from simulate disbursement.
type DisbursementAction struct {
	TargetID   string
	TargetType string
	Action     Action
}

// SimulateDisbursement to simulate disbursement status. Sandbox only.
func (c *Client) SimulateDisbursement(request SimulateDisbursementRequest) (*DisbursementAction, int, error) {
	return c.SimulateDisbursementWithContext(context.Background(), request)
}

// SimulateDisbursementWithContext to simulate disbursement status with context. Sandbox only.
func (c *Client) SimulateDisbursementWithContext(ctx context.Context, request SimulateDisbursementRequest) (*DisbursementAction, int, error) {
	if c.env == Production {
		return nil, http.StatusBadRequest, ErrSandboxOnly
	}

	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response disbursementAction
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/disbursements/%s/tasks", c.baseURL, request.ID),
		c.apiKey,
		c.secretKey,
		nil,
		request.wrap(),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return response.toDisbursementAction(), code, nil
}
