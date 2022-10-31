// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: customer.sql

package db

import (
	"context"
	"database/sql"
)

const createCustomer = `-- name: CreateCustomer :one
insert into customers(
    name,
    phone,
    address
)values(
    $1,$2,$3
) RETURNING id, phone, name, address
`

type CreateCustomerParams struct {
	Name    string         `json:"name"`
	Phone   string         `json:"phone"`
	Address sql.NullString `json:"address"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer, arg.Name, arg.Phone, arg.Address)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Name,
		&i.Address,
	)
	return i, err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
delete from customers
where id=$1
`

func (q *Queries) DeleteCustomer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, id)
	return err
}

const getCustomer = `-- name: GetCustomer :one
select id, phone, name, address from customers
where id = $1 limit 1
`

func (q *Queries) GetCustomer(ctx context.Context, id int64) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Name,
		&i.Address,
	)
	return i, err
}

const listCustomer = `-- name: ListCustomer :many
select id, phone, name, address from customers
order by name
`

func (q *Queries) ListCustomer(ctx context.Context) ([]Customer, error) {
	rows, err := q.db.QueryContext(ctx, listCustomer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.ID,
			&i.Phone,
			&i.Name,
			&i.Address,
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

const updateCustomer = `-- name: UpdateCustomer :one
update customers 
set name = $2, address = $3, phone = $4
where id = $1
returning id, phone, name, address
`

type UpdateCustomerParams struct {
	ID      int64          `json:"id"`
	Name    string         `json:"name"`
	Address sql.NullString `json:"address"`
	Phone   string         `json:"phone"`
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, updateCustomer,
		arg.ID,
		arg.Name,
		arg.Address,
		arg.Phone,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Name,
		&i.Address,
	)
	return i, err
}
