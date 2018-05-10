package irc

import (
	"fmt"
	"github.com/martin1keogh/GoCoinche/internal/models"
	bot "github.com/whyrusleeping/hellabot"
)

type Dir2Nick = map[models.Direction]string

func AddTriggers(b *bot.Bot, t *models.Table, d2n Dir2Nick) {
	triggers := []bot.Trigger{
		join(d2n),
	}

	for _, trigger := range triggers {
		b.AddTrigger(trigger)
	}
}

func command(matches func(bot.Message) bool, f func(bot.Message) (string, error)) bot.Trigger {
	trigger := bot.Trigger{
		func(b *bot.Bot, m *bot.Message) bool {
			// TODO find out why the formatting doesn't take place in .Info
			s := fmt.Sprintf("Received message %v", m.Content)
			b.Info(s)
			return matches(*m)
		},
		func(b *bot.Bot, m *bot.Message) bool {
			if !matches(*m) {
				return false
			}

			resp, err := f(*m)

			if err != nil {
				b.Reply(m, fmt.Sprintf("An error occured: %v", err))
			} else {
				b.Reply(m, resp)
			}
			return true
		},
	}

	return trigger
}
