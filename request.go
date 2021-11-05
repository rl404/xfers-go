package xfers

import (
	"net/url"
	"strconv"
	"time"
)

// ValidateBankAccountRequest is request model for validate bank account.
type ValidateBankAccountRequest struct {
	AccountNo     string   `json:"accountNo" validate:"required,numeric" mod:"no_space"`
	BankShortCode BankCode `json:"bankShortCode" validate:"required,bank_code" mod:"no_space,ucase"`
}

type validateBankAccountRequest struct {
	Data struct {
		Attributes ValidateBankAccountRequest `json:"attributes"`
	} `json:"data"`
}

func (v ValidateBankAccountRequest) wrap() validateBankAccountRequest {
	var req validateBankAccountRequest
	req.Data.Attributes = v
	return req
}

// CreateDisbursementRequest is request model for create disbursement.
type CreateDisbursementRequest struct {
	ReferenceID           string   `validate:"required" mod:"trim"`
	BankAccountHolderName string   `validate:"required" mod:"trim"`
	BankAccountNo         string   `validate:"required,numeric" mod:"no_space"`
	BankShortCode         BankCode `validate:"required,bank_code" mod:"no_space,ucase"`
	Amount                float64  `validate:"required,gt=0"`
	Description           string   `mod:"trim"`
}

type createDisbursementRequest struct {
	Data struct {
		Attributes struct {
			Amount             float64 `json:"amount"`
			ReferenceID        string  `json:"referenceId"`
			Description        string  `json:"description"`
			DisbursementMethod struct {
				Type                  string   `json:"type"`
				BankShortCode         BankCode `json:"bankShortCode"`
				BankAccountNo         string   `json:"bankAccountNo"`
				BankAccountHolderName string   `json:"bankAccountHolderName"`
			} `json:"disbursementMethod"`
		} `json:"attributes"`
	} `json:"data"`
}

func (c CreateDisbursementRequest) wrap() createDisbursementRequest {
	var r createDisbursementRequest
	r.Data.Attributes.Amount = c.Amount
	r.Data.Attributes.ReferenceID = c.ReferenceID
	r.Data.Attributes.Description = c.Description
	r.Data.Attributes.DisbursementMethod.Type = "bank_transfer"
	r.Data.Attributes.DisbursementMethod.BankShortCode = c.BankShortCode
	r.Data.Attributes.DisbursementMethod.BankAccountNo = c.BankAccountNo
	r.Data.Attributes.DisbursementMethod.BankAccountHolderName = c.BankAccountHolderName
	return r
}

// Pagination is pagination request model.
type Pagination struct {
	Page          int       `validate:"required,gte=0" mod:"default=1"`
	PageSize      int       `validate:"required,gte=0,max=1000" mod:"default=10"`
	Sort          string    `mod:"no_space"`
	CreatedAfter  time.Time ``
	CreatedBefore time.Time ``
	Status        Status    `validate:"status" mod:"no_space"`
	ReferenceID   string    `mod:"trim"`
}

func (p *Pagination) encode() string {
	query := &url.Values{}
	query.Add("page[number]", strconv.Itoa(p.Page))
	query.Add("page[size]", strconv.Itoa(p.PageSize))

	if p.Sort != "" {
		query.Add("sort", p.Sort)
	}

	if !p.CreatedAfter.IsZero() {
		query.Add("filter[createdAfter]", p.CreatedAfter.Format(time.RFC3339))
	}

	if !p.CreatedBefore.IsZero() {
		query.Add("filter[createdBefore]", p.CreatedBefore.Format(time.RFC3339))
	}

	if p.Status != "" {
		query.Add("filter[status]", string(p.Status))
	}

	if p.ReferenceID != "" {
		query.Add("filter[referenceId]", p.ReferenceID)
	}

	return query.Encode()
}

// SimulateDisbursementRequest is request model for simulate disbursement status.
type SimulateDisbursementRequest struct {
	ID     string `validate:"required" mod:"no_space"`
	Action Action `validate:"required,disbursement_action" mod:"no_space,lcase"`
}

type simulateDisbursementRequest struct {
	Data struct {
		Attributes struct {
			Action Action `json:"action"`
		} `json:"attributes"`
	} `json:"data"`
}

func (s SimulateDisbursementRequest) wrap() simulateDisbursementRequest {
	var r simulateDisbursementRequest
	r.Data.Attributes.Action = s.Action
	return r
}

// CreatePaymentRequest is request model for create payment.
type CreatePaymentRequest struct {
	PaymentMethodType        PaymentType  `validate:"required,payment_type" mod:"no_space,lcase"`
	Amount                   float64      `validate:"required,gt=0"`
	ReferenceID              string       `validate:"required" mod:"trim"`
	ExpiredAt                time.Time    `validate:"required"`
	Description              string       `mod:"trim"`
	DisplayName              string       `mod:"trim"`
	RetailOutletName         RetailOutlet `mod:"no_space,ucase"` // retail outlet
	BankShortCode            BankCode     `mod:"no_space,ucase"` // va
	SuffixNo                 string       `mod:"no_space"`       // va
	ProvideCode              EWallet      `mod:"no_space,ucase"` // e-wallet
	AfterSettlementReturnURL string       `mod:"trim"`           // e-wallet
}

type createPaymentRequest struct {
	Data struct {
		Attributes struct {
			PaymentMethodType    PaymentType `json:"paymentMethodType"`
			Amount               float64     `json:"amount"`
			ReferenceID          string      `json:"referenceId"`
			ExpiredAt            time.Time   `json:"expiredAt,omitempty"`
			Description          string      `json:"description"`
			PaymentMethodOptions struct {
				// All.
				DisplayName string `json:"displayName"`

				// Retail outlet.
				RetailOutletName RetailOutlet `json:"retailOutletName"`

				// VA.
				BankShortCode BankCode `json:"bankShortCode"`
				SuffixNo      string   `json:"suffixNo"`

				// E-wallet.
				ProviderCode             EWallet `json:"providerCode"`
				AfterSettlementReturnURL string  `json:"afterSettlementReturnUrl"`
			} `json:"paymentMethodOptions"`
		} `json:"attributes"`
	} `json:"data"`
}

type paymentRetailValidation struct {
	RetailOutletName RetailOutlet `validate:"required,retail_outlet"`
}

type paymentVAValidation struct {
	BankShortCode BankCode `validate:"required,va_bank_code"`
}

type paymentEWalletValidation struct {
	ProviderCode             EWallet `validate:"required,e_wallet"`
	AfterSettlementReturnURL string  `validate:"required,url"`
}

func (c *CreatePaymentRequest) validate() error {
	if err := validate(c); err != nil {
		return err
	}

	switch c.PaymentMethodType {
	case PaymentEWallet:
		if err := validate(&paymentEWalletValidation{ProviderCode: c.ProvideCode, AfterSettlementReturnURL: c.AfterSettlementReturnURL}); err != nil {
			return err
		}
	case PaymentOutlet:
		if err := validate(&paymentRetailValidation{RetailOutletName: c.RetailOutletName}); err != nil {
			return err
		}
	case PaymentQRIS:
		return nil
	case PaymentVA:
		if err := validate(&paymentVAValidation{BankShortCode: c.BankShortCode}); err != nil {
			return err
		}
	}

	return nil
}

func (c CreatePaymentRequest) wrap() createPaymentRequest {
	var r createPaymentRequest
	r.Data.Attributes.PaymentMethodType = c.PaymentMethodType
	r.Data.Attributes.Amount = c.Amount
	r.Data.Attributes.ReferenceID = c.ReferenceID
	r.Data.Attributes.ExpiredAt = c.ExpiredAt
	r.Data.Attributes.Description = c.Description
	r.Data.Attributes.PaymentMethodOptions.DisplayName = c.DisplayName
	r.Data.Attributes.PaymentMethodOptions.RetailOutletName = c.RetailOutletName
	r.Data.Attributes.PaymentMethodOptions.BankShortCode = c.BankShortCode
	r.Data.Attributes.PaymentMethodOptions.SuffixNo = c.SuffixNo
	r.Data.Attributes.PaymentMethodOptions.ProviderCode = c.ProvideCode
	r.Data.Attributes.PaymentMethodOptions.AfterSettlementReturnURL = c.AfterSettlementReturnURL
	return r
}

// SimulatePaymentRequest is request model for simulate payment.
type SimulatePaymentRequest struct {
	ID     string `validate:"required" mod:"no_space"`
	Action Action `validate:"required,payment_action" mod:"no_space,lcase"`
	Amount float64
}

type simulatePaymentRequest struct {
	Data struct {
		Attributes struct {
			Action Action `json:"action"`
		} `json:"attributes"`
	} `json:"data"`
}

func (s SimulatePaymentRequest) wrap() simulatePaymentRequest {
	var r simulatePaymentRequest
	r.Data.Attributes.Action = s.Action
	return r
}
