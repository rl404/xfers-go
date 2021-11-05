package xfers

import (
	"strconv"
	"time"
)

type balance struct {
	Data struct {
		Attributes Balance `json:"attributes"`
	} `json:"data"`
}

type bank struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes Bank   `json:"attributes"`
	} `json:"data"`
}

func (b *bank) toBanks() []Bank {
	banks := make([]Bank, len(b.Data))
	for i, bb := range b.Data {
		banks[i] = Bank{
			Name:      bb.Attributes.Name,
			ShortCode: bb.Attributes.ShortCode,
		}
	}
	return banks
}

type bankAccount struct {
	Data struct {
		Attributes BankAccount `json:"attributes"`
	} `json:"data"`
}

type disbursement struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			ReferenceID        string    `json:"referenceId"`
			Description        string    `json:"description"`
			Amount             string    `json:"amount"`
			Status             Status    `json:"status"`
			CreatedAt          time.Time `json:"createdAt"`
			Fees               string    `json:"fees"`
			FailureReason      string    `json:"failureReason"`
			DisbursementMethod struct {
				Type                        string   `json:"type"`
				BankAccountNo               string   `json:"bankAccountNo"`
				BankShortCode               BankCode `json:"bankShortCode"`
				BankName                    string   `json:"bankName"`
				BankAccountHolderName       string   `json:"bankAccountHolderName"`
				ServerBankAccountHolderName string   `json:"serverBankAccountHolderName"`
			} `json:"disbursementMethod"`
		} `json:"attributes"`
	} `json:"data"`
}

func (d disbursement) toDisbursement() *Disbursement {
	amount, _ := strconv.ParseFloat(d.Data.Attributes.Amount, 64)
	fees, _ := strconv.ParseFloat(d.Data.Attributes.Fees, 64)
	return &Disbursement{
		ID:                          d.Data.ID,
		ReferenceID:                 d.Data.Attributes.ReferenceID,
		Description:                 d.Data.Attributes.Description,
		Amount:                      amount,
		Status:                      d.Data.Attributes.Status,
		CreatedAt:                   d.Data.Attributes.CreatedAt,
		Fees:                        fees,
		FailureReason:               d.Data.Attributes.FailureReason,
		BankAccountNo:               d.Data.Attributes.DisbursementMethod.BankAccountNo,
		BankShortCode:               d.Data.Attributes.DisbursementMethod.BankShortCode,
		BankName:                    d.Data.Attributes.DisbursementMethod.BankName,
		BankAccountHolderName:       d.Data.Attributes.DisbursementMethod.BankAccountHolderName,
		ServerBankAccountHolderName: d.Data.Attributes.DisbursementMethod.ServerBankAccountHolderName,
	}
}

type disbursements struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			ReferenceID        string    `json:"referenceId"`
			Description        string    `json:"description"`
			Amount             string    `json:"amount"`
			Status             Status    `json:"status"`
			CreatedAt          time.Time `json:"createdAt"`
			Fees               string    `json:"fees"`
			FailureReason      string    `json:"failureReason"`
			DisbursementMethod struct {
				Type                        string   `json:"type"`
				BankAccountNo               string   `json:"bankAccountNo"`
				BankShortCode               BankCode `json:"bankShortCode"`
				BankName                    string   `json:"bankName"`
				BankAccountHolderName       string   `json:"bankAccountHolderName"`
				ServerBankAccountHolderName string   `json:"serverBankAccountHolderName"`
			} `json:"disbursementMethod"`
		} `json:"attributes"`
	} `json:"data"`
}

func (d disbursements) toDisbursements() []Disbursement {
	disbursements := make([]Disbursement, len(d.Data))
	for i, dd := range d.Data {
		amount, _ := strconv.ParseFloat(dd.Attributes.Amount, 64)
		fees, _ := strconv.ParseFloat(dd.Attributes.Fees, 64)
		disbursements[i] = Disbursement{
			ID:                          dd.ID,
			ReferenceID:                 dd.Attributes.ReferenceID,
			Description:                 dd.Attributes.Description,
			Amount:                      amount,
			Status:                      dd.Attributes.Status,
			CreatedAt:                   dd.Attributes.CreatedAt,
			Fees:                        fees,
			FailureReason:               dd.Attributes.FailureReason,
			BankAccountNo:               dd.Attributes.DisbursementMethod.BankAccountNo,
			BankShortCode:               dd.Attributes.DisbursementMethod.BankShortCode,
			BankName:                    dd.Attributes.DisbursementMethod.BankName,
			BankAccountHolderName:       dd.Attributes.DisbursementMethod.BankAccountHolderName,
			ServerBankAccountHolderName: dd.Attributes.DisbursementMethod.ServerBankAccountHolderName,
		}
	}
	return disbursements
}

type disbursementAction struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			TargetID   string `json:"targetId"`
			TargetType string `json:"targetType"`
			Action     Action `json:"action"`
		} `json:"attributes"`
	} `json:"data"`
}

func (d disbursementAction) toDisbursementAction() *DisbursementAction {
	return &DisbursementAction{
		TargetID:   d.Data.Attributes.TargetID,
		TargetType: d.Data.Attributes.TargetType,
		Action:     d.Data.Attributes.Action,
	}
}

type payment struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Status        Status    `json:"status"`
			Amount        string    `json:"amount"`
			CreatedAt     time.Time `json:"createdAt"`
			Description   string    `json:"description"`
			ExpiredAt     time.Time `json:"expiredAt"`
			ReferenceID   string    `json:"referenceId"`
			Fees          string    `json:"fees"`
			PaymentMethod struct {
				ID           string      `json:"id"`
				Type         PaymentType `json:"type"`
				ReferenceID  string      `json:"referenceId"`
				Instructions struct {
					DisplayName string `json:"displayName"`

					// Retail outlet.
					RetailOutletName RetailOutlet `json:"retailOutletName"`
					PaymentCode      string       `json:"paymentCode"`

					// VA.
					BankShortCode BankCode `json:"bankShortCode"`
					AccountNo     string   `json:"accountNo"`

					// QRIS.
					ImageURL string `json:"imageUrl"`
				} `json:"instructions"`
				// E-wallet.
				Settlement struct {
					HttpURL            string `json:"httpUrl"`
					AfterSettlementURL string `json:"afterSettlementUrl"`
				} `json:"settlement"`
			} `json:"paymentMethod"`
		} `json:"attributes"`
	} `json:"data"`
}

func (p payment) toPayment() *Payment {
	amount, _ := strconv.ParseFloat(p.Data.Attributes.Amount, 64)
	fees, _ := strconv.ParseFloat(p.Data.Attributes.Fees, 64)
	return &Payment{
		ID:                 p.Data.ID,
		Status:             p.Data.Attributes.Status,
		Amount:             amount,
		CreatedAt:          p.Data.Attributes.CreatedAt,
		Descriptions:       p.Data.Attributes.Description,
		ExpiredAt:          p.Data.Attributes.ExpiredAt,
		ReferenceID:        p.Data.Attributes.ReferenceID,
		Fees:               fees,
		PaymentMethodID:    p.Data.Attributes.PaymentMethod.ID,
		Type:               p.Data.Attributes.PaymentMethod.Type,
		DisplayName:        p.Data.Attributes.PaymentMethod.Instructions.DisplayName,
		RetailOutletName:   p.Data.Attributes.PaymentMethod.Instructions.RetailOutletName,
		PaymentCode:        p.Data.Attributes.PaymentMethod.Instructions.PaymentCode,
		BankShortCode:      p.Data.Attributes.PaymentMethod.Instructions.BankShortCode,
		AccountNo:          p.Data.Attributes.PaymentMethod.Instructions.AccountNo,
		ImageURL:           p.Data.Attributes.PaymentMethod.Instructions.ImageURL,
		HttpURL:            p.Data.Attributes.PaymentMethod.Settlement.HttpURL,
		AfterSettlementURl: p.Data.Attributes.PaymentMethod.Settlement.AfterSettlementURL,
	}
}

type payments struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Status        Status    `json:"status"`
			Amount        string    `json:"amount"`
			CreatedAt     time.Time `json:"createdAt"`
			Description   string    `json:"description"`
			ExpiredAt     time.Time `json:"expiredAt"`
			ReferenceID   string    `json:"referenceId"`
			Fees          string    `json:"fees"`
			PaymentMethod struct {
				ID           string      `json:"id"`
				Type         PaymentType `json:"type"`
				ReferenceID  string      `json:"referenceId"`
				Instructions struct {
					DisplayName string `json:"displayName"`

					// Retail outlet.
					RetailOutletName RetailOutlet `json:"retailOutletName"`
					PaymentCode      string       `json:"paymentCode"`

					// VA.
					BankShortCode BankCode `json:"bankShortCode"`
					AccountNo     string   `json:"accountNo"`

					// QRIS.
					ImageURL string `json:"imageUrl"`
				} `json:"instructions"`
				// E-wallet.
				Settlement struct {
					HttpURL            string `json:"httpUrl"`
					AfterSettlementURL string `json:"afterSettlementUrl"`
				} `json:"settlement"`
			} `json:"paymentMethod"`
		} `json:"attributes"`
	} `json:"data"`
}

func (p payments) toPayments() []Payment {
	payments := make([]Payment, len(p.Data))
	for i, pp := range p.Data {
		amount, _ := strconv.ParseFloat(pp.Attributes.Amount, 64)
		fees, _ := strconv.ParseFloat(pp.Attributes.Fees, 64)
		payments[i] = Payment{
			ID:                 pp.ID,
			Status:             pp.Attributes.Status,
			Amount:             amount,
			CreatedAt:          pp.Attributes.CreatedAt,
			Descriptions:       pp.Attributes.Description,
			ExpiredAt:          pp.Attributes.ExpiredAt,
			ReferenceID:        pp.Attributes.ReferenceID,
			Fees:               fees,
			PaymentMethodID:    pp.Attributes.PaymentMethod.ID,
			Type:               pp.Attributes.PaymentMethod.Type,
			DisplayName:        pp.Attributes.PaymentMethod.Instructions.DisplayName,
			RetailOutletName:   pp.Attributes.PaymentMethod.Instructions.RetailOutletName,
			PaymentCode:        pp.Attributes.PaymentMethod.Instructions.PaymentCode,
			BankShortCode:      pp.Attributes.PaymentMethod.Instructions.BankShortCode,
			AccountNo:          pp.Attributes.PaymentMethod.Instructions.AccountNo,
			ImageURL:           pp.Attributes.PaymentMethod.Instructions.ImageURL,
			HttpURL:            pp.Attributes.PaymentMethod.Settlement.HttpURL,
			AfterSettlementURl: pp.Attributes.PaymentMethod.Settlement.AfterSettlementURL,
		}
	}
	return payments
}

type paymentAction struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			TargetID   string `json:"targetId"`
			TargetType string `json:"targetType"`
			Action     Action `json:"action"`
		} `json:"attributes"`
	} `json:"data"`
}

func (p paymentAction) toPaymentAction() *PaymentAction {
	return &PaymentAction{
		TargetID:   p.Data.Attributes.TargetID,
		TargetType: p.Data.Attributes.TargetType,
		Action:     p.Data.Attributes.Action,
	}
}
