package database

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/vkunssec/contabius/configs"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDB *mongo.Database
)

type MongoConfig struct {
	URL       string
	AppName   string
	DebugMode bool
	Log       slog.Logger
}

func MongoDBConnection(ctx context.Context) error {
	cfg := MongoConfig{}

	cfg.URL = configs.Env("MONGOURI")
	cfg.AppName = "contabius"
	cfg.DebugMode = false

	options := options.Client().ApplyURI(cfg.URL)
	options.SetAppName(cfg.AppName)

	if configs.Env("STAGE") == "development" {
		cfg.DebugMode = true
		cfg.Log = *slog.Default()
	}

	if cfg.DebugMode {
		monitor := &event.CommandMonitor{
			Started: func(_ context.Context, e *event.CommandStartedEvent) {
				if e.CommandName != "endSessions" && e.CommandName != "ping" {
					command := e.Command.String()
					var commandJson map[string]interface{}
					err := json.Unmarshal([]byte(command), &commandJson)
					if err != nil {
						cfg.Log.Error("Error unmarshalling command", "error", err)
						return
					}
					r, _ := json.MarshalIndent(&commandJson, "", "  ")
					cfg.Log.Info(string(r))
				}
			},
			Succeeded: func(_ context.Context, e *event.CommandSucceededEvent) {
				if e.CommandName != "endSessions" && e.CommandName != "ping" {
					command := e.Reply.String()
					var commandJson map[string]interface{}
					err := json.Unmarshal([]byte(command), &commandJson)
					if err != nil {
						cfg.Log.Error("Error unmarshalling command", "error", err)
						return
					}
					r, _ := json.MarshalIndent(&commandJson, "", "  ")
					cfg.Log.Info(string(r))
				}
			},
			Failed: func(context.Context, *event.CommandFailedEvent) {},
		}
		options.SetMonitor(monitor)
	}

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return err
	}

	db := client.Database(configs.Env("MONGOURI_DATABASE"))

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	MongoDB = db
	return nil
}
