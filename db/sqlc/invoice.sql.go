// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: invoice.sql

package db

import (
	"context"
	"encoding/json"
	"time"
)

const createInvoice = `-- name: CreateInvoice :one
insert into invoices(
    customers_id,
    total_money,
    had_paid
) values(
    $1,$2,$3
) returning id, created_at, customers_id, total_money, had_paid, is_deleted
`

type CreateInvoiceParams struct {
	CustomersID int64 `json:"customers_id"`
	TotalMoney  int64 `json:"total_money"`
	HadPaid     int64 `json:"had_paid"`
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, createInvoice, arg.CustomersID, arg.TotalMoney, arg.HadPaid)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.CustomersID,
		&i.TotalMoney,
		&i.HadPaid,
		&i.IsDeleted,
	)
	return i, err
}

const createInvoiceDetail = `-- name: CreateInvoiceDetail :one
insert into invoice_detail(
    invoice_id,
    product_id,
    price_at_sell,
    amount,
    total_price,
    discount,
    last_price
)values(
    $1,$2,$3,$4,$5,$6,$7
)returning id, invoice_id, product_id, price_at_sell, amount, total_price, discount, last_price
`

type CreateInvoiceDetailParams struct {
	InvoiceID   int64   `json:"invoice_id"`
	ProductID   int64   `json:"product_id"`
	PriceAtSell int64   `json:"price_at_sell"`
	Amount      float64 `json:"amount"`
	TotalPrice  int64   `json:"total_price"`
	Discount    float64 `json:"discount"`
	LastPrice   int64   `json:"last_price"`
}

func (q *Queries) CreateInvoiceDetail(ctx context.Context, arg CreateInvoiceDetailParams) (InvoiceDetail, error) {
	row := q.db.QueryRowContext(ctx, createInvoiceDetail,
		arg.InvoiceID,
		arg.ProductID,
		arg.PriceAtSell,
		arg.Amount,
		arg.TotalPrice,
		arg.Discount,
		arg.LastPrice,
	)
	var i InvoiceDetail
	err := row.Scan(
		&i.ID,
		&i.InvoiceID,
		&i.ProductID,
		&i.PriceAtSell,
		&i.Amount,
		&i.TotalPrice,
		&i.Discount,
		&i.LastPrice,
	)
	return i, err
}

const getInvoice = `-- name: GetInvoice :one
select id, created_at, customers_id, total_money, had_paid, is_deleted from invoices
where id = $1 limit 1
`

func (q *Queries) GetInvoice(ctx context.Context, id int64) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, getInvoice, id)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.CustomersID,
		&i.TotalMoney,
		&i.HadPaid,
		&i.IsDeleted,
	)
	return i, err
}

const getInvoiceDetail = `-- name: GetInvoiceDetail :many
select id, invoice_id, product_id, price_at_sell, amount, total_price, discount, last_price from invoice_detail
where invoice_id = $1
`

func (q *Queries) GetInvoiceDetail(ctx context.Context, invoiceID int64) ([]InvoiceDetail, error) {
	rows, err := q.db.QueryContext(ctx, getInvoiceDetail, invoiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InvoiceDetail
	for rows.Next() {
		var i InvoiceDetail
		if err := rows.Scan(
			&i.ID,
			&i.InvoiceID,
			&i.ProductID,
			&i.PriceAtSell,
			&i.Amount,
			&i.TotalPrice,
			&i.Discount,
			&i.LastPrice,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listInvoice = `-- name: ListInvoice :many
select invoices.id, invoices.created_at, invoices.customers_id, invoices.total_money, invoices.had_paid, invoices.is_deleted, to_json(customers.name) as customer_name, to_json(customers.phone) as customer_phone from invoices left join customers
on invoices.customers_phone = customers.phone
`

type ListInvoiceRow struct {
	ID            int64           `json:"id"`
	CreatedAt     time.Time       `json:"created_at"`
	CustomersID   int64           `json:"customers_id"`
	TotalMoney    int64           `json:"total_money"`
	HadPaid       int64           `json:"had_paid"`
	IsDeleted     bool            `json:"is_deleted"`
	CustomerName  json.RawMessage `json:"customer_name"`
	CustomerPhone json.RawMessage `json:"customer_phone"`
}

func (q *Queries) ListInvoice(ctx context.Context) ([]ListInvoiceRow, error) {
	rows, err := q.db.QueryContext(ctx, listInvoice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListInvoiceRow
	for rows.Next() {
		var i ListInvoiceRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.CustomersID,
			&i.TotalMoney,
			&i.HadPaid,
			&i.IsDeleted,
			&i.CustomerName,
			&i.CustomerPhone,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
