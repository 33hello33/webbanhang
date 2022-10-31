package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type ProductTbl struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Unit       string  `json:"unit"`
	Amount     int64   `json:"amount"`
	Price      int64   `json:"price"`
	TotalPrice int64   `json:"total_price"`
	Discount   float64 `json:"discount"`
	LastPrice  int64   `json:"last_price"`
}

type InvoiceTxParams struct {
	CustomersPhone string       `json:"customer_phone"`
	TotalMoney     int64        `json:"total_money"`
	HadPaid        int64        `json:"had_paid"`
	Products       []ProductTbl `json:"product"`
}

type InvoiceTxResult struct {
	Invoice        Invoice         `json:"invoice"`
	InvoiceDetails []InvoiceDetail `json:"invoice_detail"`
}

func (store *Store) InvoiceTx(ctx context.Context, arg InvoiceTxParams) (InvoiceTxResult, error) {
	var result InvoiceTxResult
	var err error
	err = store.execTx(ctx, func(q *Queries) error {
		if arg.CustomersPhone == "" {
			arg.CustomersPhone = "0"
		}
		result.Invoice, err = q.CreateInvoice(ctx, CreateInvoiceParams{
			CustomersPhone: arg.CustomersPhone,
			TotalMoney:     arg.TotalMoney,
			HadPaid:        arg.HadPaid,
		})
		if err != nil {
			return err
		}

		for _, product := range arg.Products {
			invoiceDetail, err := q.CreateInvoiceDetail(ctx, CreateInvoiceDetailParams{
				InvoiceID:   result.Invoice.ID,
				ProductID:   product.ID,
				PriceAtSell: product.Price,
				Amount:      float64(product.Amount),
				TotalPrice:  product.TotalPrice,
				Discount:    product.Discount,
				LastPrice:   product.LastPrice,
			})
			if err != nil {
				return err
			}

			result.InvoiceDetails = append(result.InvoiceDetails, invoiceDetail)
		}
		return nil
	})
	return result, err
}
