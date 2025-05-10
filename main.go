package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func PrintCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	// Load Slack tokens from environment variables
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	// Validate tokens
	if botToken == "" || appToken == "" {
		log.Fatal("SLACK_BOT_TOKEN or SLACK_APP_TOKEN is not set in the environment")
	}

	// Initialize Slack bot client
	bot := slacker.NewClient(botToken, appToken)

	// Start logging command events
	go PrintCommandEvents(bot.CommandEvents())

	// Define a "ping" command
	bot.Command("ping", &slacker.CommandDefinition{
		Description: "Responds with pong",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	})

	// Start listening for events
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := bot.Listen(ctx); err != nil {
		log.Fatal(err)
	}
}
