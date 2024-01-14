package chats

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/simple-chat-app/internal/helpers"
	"github.com/wildanfaz/simple-chat-app/internal/models"
	"github.com/wildanfaz/simple-chat-app/internal/repositories"
)

type ChatsService struct {
	log      *logrus.Logger
	repo     repositories.Chats
	messages chan models.Chat
}

func NewService(log *logrus.Logger, repo repositories.Chats, messages chan models.Chat) *ChatsService {
	return &ChatsService{
		log:      log,
		repo:     repo,
		messages: messages,
	}
}

func (s *ChatsService) NewChat() func(*fiber.Ctx) error {
	return websocket.New(func(c *websocket.Conn) {
		senderId := c.Query("sender_id")
		receiverId := c.Query("receiver_id")

		senderIdInt, err := strconv.Atoi(senderId)
		if err != nil {
			c.WriteJSON(&helpers.Response{
				Error:   true,
				Message: "sender_id must be number",
			})
			return
		}

		receiverIdInt, err := strconv.Atoi(receiverId)
		if err != nil {
			c.WriteJSON(&helpers.Response{
				Error:   true,
				Message: "receiver_id must be number",
			})
			return
		}

		if senderIdInt == receiverIdInt {
			c.WriteJSON(&helpers.Response{
				Error:   true,
				Message: "sender_id with receiver_id must be different",
			})
			return
		}

		roomId := helpers.GenerateRoomID(senderIdInt, receiverIdInt)

		chats, err := s.repo.GetChats(context.Background(), roomId)
		if err != nil {
			c.WriteJSON(&helpers.Response{
				Error:   true,
				Message: err.Error(),
			})
			return
		}

		for _, chat := range chats {
			c.WriteJSON(&helpers.Response{
				Error:   false,
				Message: "success",
				Data:    chat,
			})
		}

		var (
			msg        []byte
			chat       models.Chat
			id         int
			now        = time.Now()
			lastChatId int
		)

		chat.SenderId = senderIdInt
		chat.ReceiverId = receiverIdInt
		chat.RoomId = roomId

		go func() {
			for v := range s.messages {
				if v.RoomId == chat.RoomId && lastChatId != v.ChatId {
					c.WriteJSON(&helpers.Response{
						Error:   false,
						Message: "success",
						Data:    v,
					})

					lastChatId = v.ChatId
				} else {
					lastChatId = 0
				}
			}
		}()

		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				s.log.Println("read:", err)
				break
			}
			s.log.Printf("recv: %s", msg)

			chat.Message = string(msg)
			now = time.Now()
			chat.CreatedAt = &now

			id, err = s.repo.InsertChat(context.Background(), &chat)
			if err != nil {
				s.log.Println("write:", err)
				break
			}

			chat.ChatId = id

			s.messages <- chat
			s.messages <- chat
		}
	})
}
