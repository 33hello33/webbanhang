-- name: CreateInvoice :one
insert into invoices(
    customers_phone,
    total_money,
    had_paid
) values(
    $1,$2,$3
) returning *;

-- name: GetInvoice :one
select * from invoices
where id = $1 limit 1;

-- name: ListInvoice :many
select invoices.*, to_json(customers.name) as customers_name from invoices left join customers
on invoices.customers_phone = customers.phone ;

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
select * from invoice_detail
where invoice_id = $1;