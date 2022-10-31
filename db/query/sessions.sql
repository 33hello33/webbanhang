-- name: CreateSession :one
insert into sessions(
    id,
    user_name,
    refresh_token,
    user_agent,
    client_ip,
    is_blocked,
    expired_at
) values (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSession :one
select * from sessions
where id = $1 limit 1;

-- name: DeleteSession :exec
delete from sessions
where id = $1;