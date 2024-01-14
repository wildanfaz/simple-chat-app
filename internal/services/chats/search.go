package chats

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/simple-chat-app/internal/helpers"
	"github.com/wildanfaz/simple-chat-app/internal/models"
)

func (s *ChatsService) SearchChats(c *fiber.Ctx) error {
	var (
		resp = helpers.NewResponse()
	)

	senderId := c.QueryInt("sender_id", 0)
	receiverId := c.QueryInt("receiver_id", 0)
	message := c.Query("message")

	if senderId < 1 || receiverId < 1 {
		s.log.Errorln("sender_id and receiver_id must be greater than 0")
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().
			WithMessage("sender_id and receiver_id must be greater than 0"))
	}

	if senderId != receiverId {
		s.log.Errorln("sender_id and receiver_id must be equal")
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().
			WithMessage("sender_id and receiver_id must be equal"))
	}

	if message == "" {
		s.log.Errorln("message must be not empty")
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().
			WithMessage("message must be not empty"))
	}

	chats, err := s.repo.SearchChat(c.Context(), &models.Chat{
		SenderId:   senderId,
		ReceiverId: receiverId,
		Message:    message,
	})
	if err != nil {
		s.log.Errorln("failed to search chat", err)
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().
			WithMessage("failed to search chat"))
	}

	s.log.Println("success to search chat")
	return c.JSON(resp.
		WithMessage("success to search chat").
		WithData(chats))
}
