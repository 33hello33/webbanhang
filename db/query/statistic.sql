-- name: GetTopSellingProducts :many
select to_json(products.name) as name, sum(invoice_detail.amount)  as selling_amount
from invoices left join invoice_detail
on invoices.id = invoice_detail.invoice_id
left join products 
on invoice_detail.product_id = products.id
where invoices.created_at between sqlc.arg(created_from) and sqlc.arg(created_to)
group by products.id 
order by selling_amount desc
limit sqlc.arg(number_top);
