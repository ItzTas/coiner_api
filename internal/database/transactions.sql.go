// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: transactions.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (id, sender_id, receiver_id, amount, executed_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, sender_id, receiver_id, amount, executed_at
`

type CreateTransactionParams struct {
	ID         uuid.UUID
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
	Amount     string
	ExecutedAt time.Time
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.ID,
		arg.SenderID,
		arg.ReceiverID,
		arg.Amount,
		arg.ExecutedAt,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.SenderID,
		&i.ReceiverID,
		&i.Amount,
		&i.ExecutedAt,
	)
	return i, err
}

const getUserTransactions = `-- name: GetUserTransactions :many
SELECT id, sender_id, receiver_id, amount, executed_at, 
    CASE 
        WHEN sender_id = $1 THEN 'sender'
        ELSE 'receiver'
    END AS user_role
FROM transactions
WHERE sender_id = $1 OR receiver_id = $1
`

type GetUserTransactionsRow struct {
	ID         uuid.UUID
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
	Amount     string
	ExecutedAt time.Time
	UserRole   string
}

func (q *Queries) GetUserTransactions(ctx context.Context, senderID uuid.UUID) ([]GetUserTransactionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserTransactions, senderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserTransactionsRow
	for rows.Next() {
		var i GetUserTransactionsRow
		if err := rows.Scan(
			&i.ID,
			&i.SenderID,
			&i.ReceiverID,
			&i.Amount,
			&i.ExecutedAt,
			&i.UserRole,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserTransactionsWithLimit = `-- name: GetUserTransactionsWithLimit :many
SELECT id, sender_id, receiver_id, amount, executed_at, 
    CASE 
        WHEN sender_id = $1 THEN 'sender'
        ELSE 'receiver'
    END AS user_role
FROM transactions
WHERE sender_id = $1 OR receiver_id = $1
LIMIT $2
`

type GetUserTransactionsWithLimitParams struct {
	SenderID uuid.UUID
	Limit    int64
}

type GetUserTransactionsWithLimitRow struct {
	ID         uuid.UUID
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
	Amount     string
	ExecutedAt time.Time
	UserRole   string
}

func (q *Queries) GetUserTransactionsWithLimit(ctx context.Context, arg GetUserTransactionsWithLimitParams) ([]GetUserTransactionsWithLimitRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserTransactionsWithLimit, arg.SenderID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserTransactionsWithLimitRow
	for rows.Next() {
		var i GetUserTransactionsWithLimitRow
		if err := rows.Scan(
			&i.ID,
			&i.SenderID,
			&i.ReceiverID,
			&i.Amount,
			&i.ExecutedAt,
			&i.UserRole,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
