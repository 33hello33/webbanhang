-- name: CreateAccount :one
insert into accounts(
    user_name,
    full_name,
    hash_password,
    email
)VALUES(
    $1,$2,$3,$4
)RETURNING *;

-- name: GetAccount :one
select * from accounts
where user_name=$1 limit 1;

-- name: UpdateAccount :one
update accounts
set hash_password = $2
where user_name=$1
RETURNING *;

-- name: DeleteAccount :exec
delete from accounts
where user_name=$1;