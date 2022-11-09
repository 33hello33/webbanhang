-- name: CreateSupplier :one
insert into suppliers(
    name,
    phone,
    address,
    notes
)values(
    $1,$2,$3,$4
) returning *;

-- name: GetSupplier :one
select * from suppliers
where id = $1 limit 1;

-- name: ListSupplier :many
select id, name from suppliers
order by name;

-- name: UpdateSupplier :one
update suppliers 
set name = $2, address = $3, phone = $4, notes = $5
where id = $1
returning *;

-- name: DeleteSupplier :exec
delete from suppliers
where id = $1;

-- name: SearchSupplierLikeName :many
select * from suppliers
where name like $1;

-- name: GetSupplierByPhone :one
select * from suppliers
where phone = $1;