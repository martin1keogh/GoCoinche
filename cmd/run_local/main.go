package main

import (
	"github.com/martin1keogh/GoCoinche/internal/models"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Launching run_local/main.go")
	t := models.NewTable()
	log.Info("Created table")
	t.ForeachPositionStartingFrom(models.North, func(pos *models.Position) {
		log.WithFields(log.Fields{
			"cards": pos.Hand,
		}).Info("Direction ", pos.Direction)
	})
}
