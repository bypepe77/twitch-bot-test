package bot

import (
	"github.com/gempir/go-twitch-irc/v2"
)

type Bot struct {
	Name     string
	client   *twitch.Client
	commands map[string]*Command
}

func NewBot(name string, client *twitch.Client) *Bot {
	return &Bot{
		Name:     name,
		client:   client,
		commands: make(map[string]*Command),
	}
}

func (b *Bot) Start() error {
	b.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		b.getCommandName(message)
	})

	b.client.Join("bypepe77")

	err := b.client.Connect()
	if err != nil {
		println("Error starting bot")
	}
	return nil
}

func (b *Bot) Stop() error {
	return b.client.Disconnect()
}
