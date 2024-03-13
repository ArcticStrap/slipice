package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/arcticstrap/slipice/bot/commands"
	"github.com/arcticstrap/slipice/utils/environment"
)

func main() {
	// Load environment variables
	err := environment.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	BotToken := os.Getenv("BOTTOKEN")

	// Init bot client
	bClient, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Failed to initalize bot client: %v", err)
	}

	// Register events
	bClient.AddHandler(commands.Init)

	// Listen to connection
	err = bClient.Open()
	if err != nil {
		log.Fatalf("Failed to open bot client: %v", err)
	}
	defer bClient.Close()

	// Register commands
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.Commands))
	for i, v := range commands.Commands {
		cmd, err := bClient.ApplicationCommandCreate(bClient.State.User.ID, v.GuildID, v)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	// Yield until a term signal is called
	log.Println("Bot running! Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanup
	for _, v := range registeredCommands {
		err := bClient.ApplicationCommandDelete(bClient.State.User.ID, v.GuildID, v.ID)
		if err != nil {
			log.Fatalf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
