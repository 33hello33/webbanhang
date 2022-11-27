// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: statistic.sql

package db

import (
	"context"
	"encoding/json"
	"time"
)

const getTopSellingProducts = `-- name: GetTopSellingProducts :many
select to_json(products.name) as name, sum(invoice_detail.amount)  as selling_amount
from invoices left join invoice_detail
on invoices.id = invoice_detail.invoice_id
left join products 
on invoice_detail.product_id = products.id
where invoices.created_at between $1 and $2
group by products.id 
order by selling_amount desc
limit $3
`

type GetTopSellingProductsParams struct {
	CreatedFrom time.Time `json:"created_from"`
	CreatedTo   time.Time `json:"created_to"`
	NumberTop   int32     `json:"number_top"`
}

type GetTopSellingProductsRow struct {
	Name          json.RawMessage `json:"name"`
	SellingAmount int64           `json:"selling_amount"`
}

func (q *Queries) GetTopSellingProducts(ctx context.Context, arg GetTopSellingProductsParams) ([]GetTopSellingProductsRow, error) {
	rows, err := q.db.QueryContext(ctx, getTopSellingProducts, arg.CreatedFrom, arg.CreatedTo, arg.NumberTop)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTopSellingProductsRow
	for rows.Next() {
		var i GetTopSellingProductsRow
		if err := rows.Scan(&i.Name, &i.SellingAmount); err != nil {
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