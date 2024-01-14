package routers

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/simple-chat-app/configs"
	"github.com/wildanfaz/simple-chat-app/internal/models"
	"github.com/wildanfaz/simple-chat-app/internal/pkg"
	"github.com/wildanfaz/simple-chat-app/internal/repositories"
	"github.com/wildanfaz/simple-chat-app/internal/services/chats"
)

func InitFiber() {
	app := fiber.New()

	db := configs.InitPostgreSQL()
	repo := repositories.NewChats(db)
	log := pkg.InitLogger()
	messages := make(chan models.Chat, 255)

	chatsService := chats.NewService(log, repo, messages)

	app.Get("/chats", chatsService.SearchChats)

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", chatsService.NewChat())

	log.Fatal(app.Listen(":3000"))
}
