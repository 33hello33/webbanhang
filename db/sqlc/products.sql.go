// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: products.sql

package db

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :one
insert into products(
    name,
    unit,
    price,
    price_import,
    amount,
    warehouse,
    id_supplier
)values(
    $1,$2,$3,$4,$5,$6,$7
) returning id, name, unit, price_import, amount, price, warehouse, created_at, id_supplier
`

type CreateProductParams struct {
	Name        string        `json:"name"`
	Unit        string        `json:"unit"`
	Price       int64         `json:"price"`
	PriceImport int64         `json:"price_import"`
	Amount      int64         `json:"amount"`
	Warehouse   string        `json:"warehouse"`
	IDSupplier  sql.NullInt64 `json:"id_supplier"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.Name,
		arg.Unit,
		arg.Price,
		arg.PriceImport,
		arg.Amount,
		arg.Warehouse,
		arg.IDSupplier,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Unit,
		&i.PriceImport,
		&i.Amount,
		&i.Price,
		&i.Warehouse,
		&i.CreatedAt,
		&i.IDSupplier,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
delete from products where id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
select id, name, unit, price_import, amount, price, warehouse, created_at, id_supplier from products
where id=$1 limit 1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Unit,
		&i.PriceImport,
		&i.Amount,
		&i.Price,
		&i.Warehouse,
		&i.CreatedAt,
		&i.IDSupplier,
	)
	return i, err
}

const getProductForUpdate = `-- name: GetProductForUpdate :one
select id, name, unit, price_import, amount, price, warehouse, created_at, id_supplier from products
where id=$1 limit 1
for no key update
`

func (q *Queries) GetProductForUpdate(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductForUpdate, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Unit,
		&i.PriceImport,
		&i.Amount,
		&i.Price,
		&i.Warehouse,
		&i.CreatedAt,
		&i.IDSupplier,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
select id, name from products
order by name
`

type ListProductsRow struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) ListProducts(ctx context.Context) ([]ListProductsRow, error) {
	rows, err := q.db.QueryContext(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListProductsRow
	for rows.Next() {
		var i ListProductsRow
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

const searchProductLikeName = `-- name: SearchProductLikeName :many
select id, name, unit, price_import, amount, price, warehouse, created_at, id_supplier from products
where name like $1
`

func (q *Queries) SearchProductLikeName(ctx context.Context, name string) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, searchProductLikeName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Unit,
			&i.PriceImport,
			&i.Amount,
			&i.Price,
			&i.Warehouse,
			&i.CreatedAt,
			&i.IDSupplier,
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

const updateProduct = `-- name: UpdateProduct :one
update products
set amount = $2, 
price = $3, 
price_import = $4, 
warehouse = $5,
id_supplier = $6,
unit = $7,
name = $8
where id = $1
returning id, name, unit, price_import, amount, price, warehouse, created_at, id_supplier
`

type UpdateProductParams struct {
	ID          int64         `json:"id"`
	Amount      int64         `json:"amount"`
	Price       int64         `json:"price"`
	PriceImport int64         `json:"price_import"`
	Warehouse   string        `json:"warehouse"`
	IDSupplier  sql.NullInt64 `json:"id_supplier"`
	Unit        string        `json:"unit"`
	Name        string        `json:"name"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.ID,
		arg.Amount,
		arg.Price,
		arg.PriceImport,
		arg.Warehouse,
		arg.IDSupplier,
		arg.Unit,
		arg.Name,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Unit,
		&i.PriceImport,
		&i.Amount,
		&i.Price,
		&i.Warehouse,
		&i.CreatedAt,
		&i.IDSupplier,
	)
	return i, err
}
