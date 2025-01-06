-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currency) VALUES ($1, $2, $3) RETURNING *;

-- name: CreateUser :one
INSERT INTO users (username, pass, full_name, email) VALUES ($1, $2, $3, $4) RETURNING *;


-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;


-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id LIMIT $2 OFFSET $1;

-- name: UpdateAccount :one
UPDATE accounts SET balance= $2 WHERE id = $1 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;
-- Create a transfer record
-- name: CreateTransfer :one
INSERT INTO transfers (
    from_acc_id, 
    to_acc_id, 
    amount
) VALUES (
    $1, $2, $3
) RETURNING id, from_acc_id, to_acc_id, amount, created_at;

-- Create a ledger entry
-- name: CreateEntry :one
INSERT INTO entries (
    acc_id, 
    amount
) VALUES (
    $1, $2
) RETURNING id, acc_id, amount, created_at;

-- Fetch account details
-- name: GetAccountForUpdate :one
SELECT id, owner, balance, currency, created_at 
FROM accounts 
WHERE id = $1 
FOR UPDATE;

-- Update an account balance
-- name: UpdateAccountBalance :one
UPDATE accounts 
SET balance = balance + $2 
WHERE id = $1 RETURNING *;

-- name: GetTransfer :one
SELECT 
    id, 
    from_acc_id, 
    to_acc_id, 
    amount, 
    created_at 
FROM transfers 
WHERE id = $1 
LIMIT 1;
