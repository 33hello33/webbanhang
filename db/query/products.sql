-- name: CreateProduct :one
insert into products(
    name,
    unit,
    price,
    price_import,
    amount,
    warehouse
)values(
    $1,$2,$3,$4,$5,$6
) returning *;

-- name: GetProduct :one
select * from products
where id=$1 limit 1;

-- name: GetProductForUpdate :one
select * from products
where id=$1 limit 1
for no key update;

-- name: ListProducts :many
select id, name from products
order by name;

-- name: UpdateProduct :one
update products
set amount = $2, 
price = $3, 
price_import = $4, 
warehouse = $5,
unit = $6,
name = $7
where id = $1
returning *;

-- name: DeleteProduct :exec
delete from products where id = $1;


-- name: SearchProductLikeName :many
select * from products
where name like $1;

-- name: UpdateAmountProduct :one
update products
set amount=amount-$2
where id=$1
returning *;

-- name: CopyProduct :one
insert into products(
    name,
    unit,
    price,
    price_import,
    amount,
    warehouse)
select concat(name, '(copy)'), 
    unit,
    price,
    price_import,
    amount,
    warehouse
from products as pd
where pd.id = $1
returning *;