package database

import (
	"contabius/configs"
	"context"
	"encoding/json"
	"log/slog"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDB *mongo.Database
)

type MongoConfig struct {
	URL       string      // URL de conexão com o MongoDB
	AppName   string      // Nome da aplicação
	DebugMode bool        // Flag que habilita os logs de debug
	Log       slog.Logger // Logger que será utilizado
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
					json.Unmarshal([]byte(command), &commandJson)
					r, _ := json.MarshalIndent(&commandJson, "", "  ")

					cfg.Log.Info(string(r))
				}
			},
			Succeeded: func(_ context.Context, e *event.CommandSucceededEvent) {
				if e.CommandName != "endSessions" && e.CommandName != "ping" {
					command := e.Reply.String()

					var commandJson map[string]interface{}
					json.Unmarshal([]byte(command), &commandJson)
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
