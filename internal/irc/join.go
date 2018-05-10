package irc

import (
	"fmt"
	"github.com/martin1keogh/GoCoinche/internal/models"
	bot "github.com/whyrusleeping/hellabot"
)

func join(d2n Dir2Nick) bot.Trigger {
	return command(
		func(m bot.Message) bool { return m.Content == "!join" },
		func(m bot.Message) (string, error) {
			// TODO implement 'table' side
			return registerPlayer(m.From, d2n)
		})
}

func registerPlayer(nick string, d2n Dir2Nick) (string, error) {
	if len(d2n) >= 4 {
		return "", fmt.Errorf("Table is full")
	}

	for _, player := range d2n {
		if player == nick {
			return "", fmt.Errorf("Player %s is already playing at this table", nick)
		}
	}

	for _, dir := range models.AllDirections() {
		_, inUse := d2n[dir]
		if !inUse {
			d2n[dir] = nick
			return fmt.Sprintf("%s joins the table at position %v", nick, dir), nil
		}
	}

	return "", fmt.Errorf("Unkown error: unable to register %s in %v.", nick, d2n)
}
