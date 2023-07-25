-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByEmailAndPassword :one
SELECT (password = crypt($1, (SELECT password FROM users WHERE users.email=$2))) AS password_matches,
       email, username, firstname, lastname FROM users WHERE email=$2 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserByUsernameAndPassword :one
SELECT (password = crypt($1, (SELECT password FROM users WHERE users.username=$2))) AS password_matches,
       email, username, firstname, lastname FROM users WHERE username=$2 LIMIT 1;

-- name: ListUsers :many
SELECT id, username, firstname, lastname, email FROM users
ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (username, firstname, lastname, email, password)
VALUES ($1, $2, $3, $4, crypt($5, gen_salt('bf'))) RETURNING id;

-- name: AssignUserRole :exec
INSERT INTO user_roles (
    user_id, roles_id
) VALUES ( $1, (SELECT id FROM roles WHERE name = $2) );

-- name: GetRoleByUsername :one
SELECT name FROM roles WHERE id=(
    SELECT roles_id FROM user_roles WHERE user_id=
    (
        SELECT id FROM users WHERE username=$1
    )
);