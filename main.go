package main

import (
	"os"

	server "github.com/AmirhoseinMasoumi/notification-service/api"
	"github.com/AmirhoseinMasoumi/notification-service/internal/channels"
	"github.com/AmirhoseinMasoumi/notification-service/internal/queue"
	config "github.com/AmirhoseinMasoumi/notification-service/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	emailChannel := channels.NewEmailChannel(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUsername, cfg.SMTPPassword)
	htmlEmailChannel := channels.NewHtmlEmailChannel(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUsername, cfg.SMTPPassword)
	smsChannel := channels.NewSMSChannel(cfg.SMSAPIKey)
	pushChannel := channels.NewPushChannel(cfg.PushAPIKey)

	jobQueue := queue.NewJobQueue(cfg.JobQueueSize)
	for i := 0; i < cfg.JobWorkerCount; i++ {
		jobQueue.StartWorker(i)
	}

	srv := server.NewServer(emailChannel, htmlEmailChannel, smsChannel, pushChannel, jobQueue)

	log.Info().Msg("Starting server on: " + cfg.ServerAddress)
	err := srv.Router.Run(cfg.ServerAddress)
	if err != nil {
		log.Fatal().Err(err)
	}
}
