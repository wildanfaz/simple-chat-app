package main

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wildanfaz/simple-chat-app/cmd"
)

func main() {
	cmd.InitCmd(context.Background())
}
