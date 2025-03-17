package nats

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log/slog"
	"os"
	"time"
)

type LogMessage struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

func setupNatsLogger(natsURL string) (*slog.logger, *nats.Conn) {

	nc, err := nats.Connect(natsURL)
	if err != nil {
		panic(err)
	}

	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)

	log = slog.New(
		slog.HandlerFunc(
			func(r slog.Record) error {

				msg := LogMessage{
					Level:   r.Level.String(),
					Message: r.Message,
					Time:    time.Now().Format(time.RFC3339),
				}

				data, err := json.Marshal(msg)
				if err != nil {
					return err
				}

				return nc.Publish("logs", data)

			},
		),
	)
	return log, nc
}
