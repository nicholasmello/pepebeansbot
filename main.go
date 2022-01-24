package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Getting bot token from -t flag
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token
	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	bot.AddHandler(messageCreate)

	bot.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening
	err = bot.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal,1)
	signal.Notify(sc,syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
	fmt.Println("")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore message if it is created by self
	if m.Author.ID == s.State.User.ID {return}

	if m.ChannelID == "849436421382996000" {
		s.MessageReactionAdd(m.ChannelID, m.ID, "3542pepebean:905286131380289646")
	}
}