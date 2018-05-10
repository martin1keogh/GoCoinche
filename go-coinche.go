package main

import (
	"github.com/martin1keogh/GoCoinche/internal/irc"
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

	irc.AddTriggers(cb, &table, direction2Nick)

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
