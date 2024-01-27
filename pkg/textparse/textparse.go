package textparse

import "github.com/bwmarrin/discordgo"

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  // Ignore bot messages
  if m.Author.ID == s.State.User.ID {
    return
  }
}
