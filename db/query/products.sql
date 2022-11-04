-- name: CreateProduct :one
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
id_supplier = $6,
unit = $7,
name = $8
where id = $1
returning *;

-- name: DeleteProduct :exec
delete from products where id = $1;


-- name: SearchProductLikeName :many
select * from products
where name like $1;