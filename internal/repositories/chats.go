package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/wildanfaz/simple-chat-app/internal/models"
)

type ImplChats struct {
	db *pgx.Conn
}

type Chats interface {
	GetChats(ctx context.Context, roomId string) (models.Chats, error)
	InsertChat(ctx context.Context, chat *models.Chat) (int, error)
	SearchChat(ctx context.Context, chat *models.Chat) (models.Chats, error)
}

func NewChats(db *pgx.Conn) Chats {
	return &ImplChats{
		db: db,
	}
}

func (r *ImplChats) GetChats(ctx context.Context, roomId string) (models.Chats, error) {
	var (
		chats models.Chats
		chat  models.Chat
	)

	rows, err := r.db.Query(ctx, `
	SELECT chat_id, room_id, sender_id, receiver_id, message, created_at, updated_at 
	FROM chats 
	WHERE room_id = $1
	`, roomId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&chat.ChatId, &chat.RoomId, &chat.SenderId, &chat.ReceiverId, &chat.Message, &chat.CreatedAt, &chat.UpdatedAt)
		if err != nil {
			return nil, err
		}

		chat.ToLocal()
		chats = append(chats, chat)
	}

	return chats, nil
}

func (r *ImplChats) InsertChat(ctx context.Context, chat *models.Chat) (int, error) {
	var (
		id int
	)

	err := r.db.QueryRow(ctx, `
	INSERT INTO chats (room_id, sender_id, receiver_id, message, created_at) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING chat_id
	`, chat.RoomId, chat.SenderId, chat.ReceiverId, chat.Message, chat.CreatedAt).Scan(&id)

	return id, err
}

func (r *ImplChats) SearchChat(ctx context.Context, chat *models.Chat) (models.Chats, error) {
	var (
		chats  models.Chats
		result models.Chat
	)

	chat.Message = "%" + chat.Message + "%"

	rows, err := r.db.Query(ctx, `
	SELECT chat_id, room_id, sender_id, receiver_id, message, created_at, updated_at
	FROM chats
	WHERE (sender_id = $1 OR receiver_id = $2)
	AND message ILIKE $3
	`, chat.SenderId, chat.ReceiverId, chat.Message)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&result.ChatId, &result.RoomId, &result.SenderId, &result.ReceiverId, &result.Message, &result.CreatedAt, &result.UpdatedAt)
		if err != nil {
			return nil, err
		}

		result.ToLocal()
		chats = append(chats, result)
	}

	return chats, nil
}
