package main

import (
  "log"
  "os"

  "github.com/bwmarrin/discordgo"
  "github.com/joho/godotenv"
)

func main() {
  // Load environment variables
  err := godotenv.Load()
  if err != nil {
    log.Fatalf("Failed to load .env file: %v", err)
  }

  BotToken := os.Getenv("BOTTOKEN")

  // Init bot client
  bClient, err := discordgo.New("Bot " + BotToken)
  if err != nil {
    log.Fatalf("Failed to initalize bot client: %v", err)
  }

  err = bClient.Open()
  if err != nil {
    log.Fatalf("Failed to open bot client: %v", err)
  }
  defer bClient.Close()
}
