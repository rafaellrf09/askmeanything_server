// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package pgstore

import (
	"context"

	"github.com/google/uuid"
)

const getMessage = `-- name: GetMessage :one
SELECT
     "id", "room_id", "message", "reactions_count", "answered"
FROM messages
WHERE room_id = $1
`

func (q *Queries) GetMessage(ctx context.Context, roomID uuid.UUID) (Message, error) {
	row := q.db.QueryRow(ctx, getMessage, roomID)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.Message,
		&i.ReactionsCount,
		&i.Answered,
	)
	return i, err
}

const getRoom = `-- name: GetRoom :one
SELECT
    "id", "theme"
FROM rooms
WHERE id = $1
`

func (q *Queries) GetRoom(ctx context.Context, id uuid.UUID) (Room, error) {
	row := q.db.QueryRow(ctx, getRoom, id)
	var i Room
	err := row.Scan(&i.ID, &i.Theme)
	return i, err
}

const getRooms = `-- name: GetRooms :many
SELECT
FROM rooms
`

type GetRoomsRow struct {
}

func (q *Queries) GetRooms(ctx context.Context) ([]GetRoomsRow, error) {
	rows, err := q.db.Query(ctx, getRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRoomsRow
	for rows.Next() {
		var i GetRoomsRow
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertMessage = `-- name: InsertMessage :one
INSERT INTO messages
    ("room_id", "message") VALUES
    ($1 , $2)
RETURNING "id"
`

type InsertMessageParams struct {
	RoomID  uuid.UUID
	Message string
}

func (q *Queries) InsertMessage(ctx context.Context, arg InsertMessageParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertMessage, arg.RoomID, arg.Message)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const insertRoom = `-- name: InsertRoom :one
INSERT INTO rooms
    ("theme") VALUES
    ( $1 )
RETURNING "id"
`

func (q *Queries) InsertRoom(ctx context.Context, theme string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertRoom, theme)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const markMessageAsAnswered = `-- name: MarkMessageAsAnswered :exec
UPDATE messages
SET
    answered = true
WHERE
    id = $1
`

func (q *Queries) MarkMessageAsAnswered(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, markMessageAsAnswered, id)
	return err
}

const reactToMessage = `-- name: ReactToMessage :one
UPDATE messages
SET
    reactions_count = reactions_count + 1
WHERE id = $1
RETURNING reactions_count
`

func (q *Queries) ReactToMessage(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, reactToMessage, id)
	var reactions_count int64
	err := row.Scan(&reactions_count)
	return reactions_count, err
}

const removeReactionFromMessage = `-- name: RemoveReactionFromMessage :one
UPDATE messages
SET
    reactions_count = reactions_count - 1
WHERE id = $1
RETURNING reactions_count
`

func (q *Queries) RemoveReactionFromMessage(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, removeReactionFromMessage, id)
	var reactions_count int64
	err := row.Scan(&reactions_count)
	return reactions_count, err
}