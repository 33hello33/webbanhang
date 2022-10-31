-- name: CreateCustomer :one
insert into customers(
    name,
    phone,
    address
)values(
    $1,$2,$3
) RETURNING *;

-- name: ListCustomer :many
select * from customers
order by name;

-- name: DeleteCustomer :exec
delete from customers
where id=$1;

-- name: GetCustomer :one
select * from customers
where id = $1 limit 1;

-- name: UpdateCustomer :one
update customers 
set name = $2, address = $3, phone = $4
where id = $1
returning *;
