-- name: CreateInvoice :one
insert into invoices(
    customers_id,
    total_money,
    had_paid,
    is_done
) values(
    $1,$2,$3,$4
) returning *;

-- name: GetInvoice :one
select * from invoices
where id = $1 limit 1;

-- name: ListInvoice :many
select invoices.*, to_json(customers.name) as customer_name, to_json(customers.phone) as customer_phone from invoices left join customers
on invoices.customers_id = customers.id ;

-- name: CreateInvoiceDetail :one
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
)returning *;

-- name: GetInvoiceDetail :many
select invoice_detail.*, to_json(products.name) as product_name,  to_json(products.unit) as product_unit
from invoice_detail left join products
    on invoice_detail.product_id = products.id
where invoice_id = $1;

-- name: FindInvoice :many
select invoices.*, to_json(name) as name, to_json(phone) as phone 
from invoices left join customers
    on invoices.customers_id = customers.id 
where (created_at between sqlc.arg(created_from) and sqlc.arg(created_to))
    and (name like coalesce(sqlc.narg(name), name)) 
    and (invoices.id = coalesce(sqlc.narg(id_invoice), invoices.id))
    and (invoices.is_done = coalesce(sqlc.narg(is_done),is_done)); 

-- name: SumToTalMoney :one
select sum(total_money)
from invoices left join customers
    on invoices.customers_id = customers.id 
where (created_at between sqlc.arg(created_from) and sqlc.arg(created_to))
    and (name like coalesce(sqlc.narg(name), name)) 
    and (invoices.id = coalesce(sqlc.narg(id_invoice), invoices.id))
    and (invoices.is_done = coalesce(sqlc.narg(is_done),is_done)); 
