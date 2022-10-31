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
where phone=$1;

-- name: GetCustomer :one
select * from customers
where phone = $1 limit 1;
