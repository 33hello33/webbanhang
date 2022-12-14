// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: supplier.sql

package db

import (
	"context"
	"database/sql"
)

const createSupplier = `-- name: CreateSupplier :one
insert into suppliers(
    name,
    phone,
    zalo,
    address,
    notes,
    bank_name,
    bank_number
)values(
    $1,$2,$3,$4,$5,$6,$7
) returning id, name, phone, zalo, address, notes, bank_name, bank_number
`

type CreateSupplierParams struct {
	Name       string         `json:"name"`
	Phone      string         `json:"phone"`
	Zalo       sql.NullString `json:"zalo"`
	Address    sql.NullString `json:"address"`
	Notes      sql.NullString `json:"notes"`
	BankName   sql.NullString `json:"bank_name"`
	BankNumber sql.NullString `json:"bank_number"`
}

func (q *Queries) CreateSupplier(ctx context.Context, arg CreateSupplierParams) (Supplier, error) {
	row := q.db.QueryRowContext(ctx, createSupplier,
		arg.Name,
		arg.Phone,
		arg.Zalo,
		arg.Address,
		arg.Notes,
		arg.BankName,
		arg.BankNumber,
	)
	var i Supplier
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.Zalo,
		&i.Address,
		&i.Notes,
		&i.BankName,
		&i.BankNumber,
	)
	return i, err
}

const deleteSupplier = `-- name: DeleteSupplier :exec
delete from suppliers
where id = $1
`

func (q *Queries) DeleteSupplier(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSupplier, id)
	return err
}

const getSupplier = `-- name: GetSupplier :one
select id, name, phone, zalo, address, notes, bank_name, bank_number from suppliers
where id = $1 limit 1
`

func (q *Queries) GetSupplier(ctx context.Context, id int64) (Supplier, error) {
	row := q.db.QueryRowContext(ctx, getSupplier, id)
	var i Supplier
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.Zalo,
		&i.Address,
		&i.Notes,
		&i.BankName,
		&i.BankNumber,
	)
	return i, err
}

const getSupplierByPhone = `-- name: GetSupplierByPhone :one
select id, name, phone, zalo, address, notes, bank_name, bank_number from suppliers
where phone = $1
`

func (q *Queries) GetSupplierByPhone(ctx context.Context, phone string) (Supplier, error) {
	row := q.db.QueryRowContext(ctx, getSupplierByPhone, phone)
	var i Supplier
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.Zalo,
		&i.Address,
		&i.Notes,
		&i.BankName,
		&i.BankNumber,
	)
	return i, err
}

const listSupplier = `-- name: ListSupplier :many
select id, name from suppliers
order by name
`

type ListSupplierRow struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) ListSupplier(ctx context.Context) ([]ListSupplierRow, error) {
	rows, err := q.db.QueryContext(ctx, listSupplier)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSupplierRow
	for rows.Next() {
		var i ListSupplierRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const searchSupplierLikeName = `-- name: SearchSupplierLikeName :many
select id, name, phone, zalo, address, notes, bank_name, bank_number from suppliers
where name like $1
`

func (q *Queries) SearchSupplierLikeName(ctx context.Context, name string) ([]Supplier, error) {
	rows, err := q.db.QueryContext(ctx, searchSupplierLikeName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Supplier
	for rows.Next() {
		var i Supplier
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Phone,
			&i.Zalo,
			&i.Address,
			&i.Notes,
			&i.BankName,
			&i.BankNumber,
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

const updateSupplier = `-- name: UpdateSupplier :one
update suppliers 
set name = $2, address = $3, phone = $4, zalo = $5, notes = $6, bank_name=$7, bank_number=$8
where id = $1
returning id, name, phone, zalo, address, notes, bank_name, bank_number
`

type UpdateSupplierParams struct {
	ID         int64          `json:"id"`
	Name       string         `json:"name"`
	Address    sql.NullString `json:"address"`
	Phone      string         `json:"phone"`
	Zalo       sql.NullString `json:"zalo"`
	Notes      sql.NullString `json:"notes"`
	BankName   sql.NullString `json:"bank_name"`
	BankNumber sql.NullString `json:"bank_number"`
}

func (q *Queries) UpdateSupplier(ctx context.Context, arg UpdateSupplierParams) (Supplier, error) {
	row := q.db.QueryRowContext(ctx, updateSupplier,
		arg.ID,
		arg.Name,
		arg.Address,
		arg.Phone,
		arg.Zalo,
		arg.Notes,
		arg.BankName,
		arg.BankNumber,
	)
	var i Supplier
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.Zalo,
		&i.Address,
		&i.Notes,
		&i.BankName,
		&i.BankNumber,
	)
	return i, err
}
