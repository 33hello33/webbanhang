// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: invoice.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const createInvoice = `-- name: CreateInvoice :one
insert into invoices(
    customers_id,
    total_money,
    had_paid,
    is_done
) values(
    $1,$2,$3,$4
) returning id, created_at, customers_id, total_money, had_paid, is_done, is_deleted
`

type CreateInvoiceParams struct {
	CustomersID int64 `json:"customers_id"`
	TotalMoney  int64 `json:"total_money"`
	HadPaid     int64 `json:"had_paid"`
	IsDone      bool  `json:"is_done"`
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, createInvoice,
		arg.CustomersID,
		arg.TotalMoney,
		arg.HadPaid,
		arg.IsDone,
	)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.CustomersID,
		&i.TotalMoney,
		&i.HadPaid,
		&i.IsDone,
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

const findInvoice = `-- name: FindInvoice :many
select invoices.id, invoices.created_at, invoices.customers_id, invoices.total_money, invoices.had_paid, invoices.is_done, invoices.is_deleted, to_json(name) as name, to_json(phone) as phone 
from invoices left join customers
    on invoices.customers_id = customers.id 
where (created_at between $1 and $2)
    and (name like coalesce($3, name)) 
    and (invoices.id = coalesce($4, invoices.id))
    and (invoices.is_done = coalesce($5,is_done))
`

type FindInvoiceParams struct {
	CreatedFrom time.Time      `json:"created_from"`
	CreatedTo   time.Time      `json:"created_to"`
	Name        sql.NullString `json:"name"`
	IDInvoice   sql.NullInt64  `json:"id_invoice"`
	IsDone      sql.NullBool   `json:"is_done"`
}

type FindInvoiceRow struct {
	ID          int64           `json:"id"`
	CreatedAt   time.Time       `json:"created_at"`
	CustomersID int64           `json:"customers_id"`
	TotalMoney  int64           `json:"total_money"`
	HadPaid     int64           `json:"had_paid"`
	IsDone      bool            `json:"is_done"`
	IsDeleted   bool            `json:"is_deleted"`
	Name        json.RawMessage `json:"name"`
	Phone       json.RawMessage `json:"phone"`
}

func (q *Queries) FindInvoice(ctx context.Context, arg FindInvoiceParams) ([]FindInvoiceRow, error) {
	rows, err := q.db.QueryContext(ctx, findInvoice,
		arg.CreatedFrom,
		arg.CreatedTo,
		arg.Name,
		arg.IDInvoice,
		arg.IsDone,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindInvoiceRow
	for rows.Next() {
		var i FindInvoiceRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.CustomersID,
			&i.TotalMoney,
			&i.HadPaid,
			&i.IsDone,
			&i.IsDeleted,
			&i.Name,
			&i.Phone,
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

const getInvoice = `-- name: GetInvoice :one
select id, created_at, customers_id, total_money, had_paid, is_done, is_deleted from invoices
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
		&i.IsDone,
		&i.IsDeleted,
	)
	return i, err
}

const getInvoiceDetail = `-- name: GetInvoiceDetail :many
select invoice_detail.id, invoice_detail.invoice_id, invoice_detail.product_id, invoice_detail.price_at_sell, invoice_detail.amount, invoice_detail.total_price, invoice_detail.discount, invoice_detail.last_price, to_json(products.name) as product_name,  to_json(products.unit) as product_unit
from invoice_detail left join products
    on invoice_detail.product_id = products.id
where invoice_id = $1
`

type GetInvoiceDetailRow struct {
	ID          int64           `json:"id"`
	InvoiceID   int64           `json:"invoice_id"`
	ProductID   int64           `json:"product_id"`
	PriceAtSell int64           `json:"price_at_sell"`
	Amount      float64         `json:"amount"`
	TotalPrice  int64           `json:"total_price"`
	Discount    float64         `json:"discount"`
	LastPrice   int64           `json:"last_price"`
	ProductName json.RawMessage `json:"product_name"`
	ProductUnit json.RawMessage `json:"product_unit"`
}

func (q *Queries) GetInvoiceDetail(ctx context.Context, invoiceID int64) ([]GetInvoiceDetailRow, error) {
	rows, err := q.db.QueryContext(ctx, getInvoiceDetail, invoiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetInvoiceDetailRow
	for rows.Next() {
		var i GetInvoiceDetailRow
		if err := rows.Scan(
			&i.ID,
			&i.InvoiceID,
			&i.ProductID,
			&i.PriceAtSell,
			&i.Amount,
			&i.TotalPrice,
			&i.Discount,
			&i.LastPrice,
			&i.ProductName,
			&i.ProductUnit,
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
select invoices.id, invoices.created_at, invoices.customers_id, invoices.total_money, invoices.had_paid, invoices.is_done, invoices.is_deleted, to_json(customers.name) as customer_name, to_json(customers.phone) as customer_phone from invoices left join customers
on invoices.customers_id = customers.id
`

type ListInvoiceRow struct {
	ID            int64           `json:"id"`
	CreatedAt     time.Time       `json:"created_at"`
	CustomersID   int64           `json:"customers_id"`
	TotalMoney    int64           `json:"total_money"`
	HadPaid       int64           `json:"had_paid"`
	IsDone        bool            `json:"is_done"`
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
			&i.IsDone,
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

const sumToTalMoney = `-- name: SumToTalMoney :one
select sum(total_money)
from invoices left join customers
    on invoices.customers_id = customers.id 
where (created_at between $1 and $2)
    and (name like coalesce($3, name)) 
    and (invoices.id = coalesce($4, invoices.id))
    and (invoices.is_done = coalesce($5,is_done))
`

type SumToTalMoneyParams struct {
	CreatedFrom time.Time      `json:"created_from"`
	CreatedTo   time.Time      `json:"created_to"`
	Name        sql.NullString `json:"name"`
	IDInvoice   sql.NullInt64  `json:"id_invoice"`
	IsDone      sql.NullBool   `json:"is_done"`
}

func (q *Queries) SumToTalMoney(ctx context.Context, arg SumToTalMoneyParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, sumToTalMoney,
		arg.CreatedFrom,
		arg.CreatedTo,
		arg.Name,
		arg.IDInvoice,
		arg.IsDone,
	)
	var sum int64
	err := row.Scan(&sum)
	return sum, err
}

const updateTotalMoneyInvoice = `-- name: UpdateTotalMoneyInvoice :exec
update invoices
set had_paid = total_money,
    is_done = true,
    created_at = now()
where id = $1
`

func (q *Queries) UpdateTotalMoneyInvoice(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, updateTotalMoneyInvoice, id)
	return err
}
