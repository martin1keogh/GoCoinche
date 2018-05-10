package main

import (
	"fmt"

	"github.com/martin1keogh/GoCoinche/internal/models"
	log "github.com/sirupsen/logrus"
	bot "github.com/whyrusleeping/hellabot"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

var logger = log.New()

type Dir2Nick = map[models.Direction]string

func main() {
	table := models.NewTable()
	direction2Nick := make(Dir2Nick)

	// TODO pass by parameters
	cb, err := bot.NewBot("localhost:6667", "cb")
	if err != nil {
		panic(err)
	}

	cb.Logger.SetHandler(WrappedLogrus{*logger})

	addTriggers(cb, &table, direction2Nick)

	cb.Run()
}

// Good to know logging sucks anywhere you go
type WrappedLogrus struct {
	underlying log.Logger
}

func (wl WrappedLogrus) Log(r *(log15.Record)) error {
	// TODO find a way to add the context to the log
	// wf := wl.WithFields(log.Fields(r.Ctx))
	switch r.Lvl {
	case log15.LvlDebug:
		wl.underlying.Debug(r.Msg)
	case log15.LvlInfo:
		wl.underlying.Info(r.Msg)
	case log15.LvlWarn:
		wl.underlying.Warn(r.Msg)
	case log15.LvlError:
		wl.underlying.Error(r.Msg)
	default:
		wl.underlying.Warn("Unhandled log15 log level %v, logging as warn", r.Lvl)
		wl.underlying.Warn(r.Msg)
	}
	return nil
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

// TODO move this to its own package
// TODO move each of the commands to their own file
func addTriggers(b *bot.Bot, t *models.Table, d2n Dir2Nick) {
	triggers := []bot.Trigger{
		// !join
		command(
			func(m bot.Message) bool { return m.Content == "!join" },
			func(m bot.Message) (string, error) {
				// TODO implement 'table' side
				return registerPlayer(m.From, d2n)
			}),
	}

	for _, trigger := range triggers {
		b.AddTrigger(trigger)
	}
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
