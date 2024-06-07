package server

import (
	"net/http"
	"time"

	"github.com/AmirhoseinMasoumi/notification-service/internal/channels"
	"github.com/AmirhoseinMasoumi/notification-service/internal/queue"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Server struct {
	EmailChannel     *channels.EmailChannel
	HtmlEmailChannel *channels.HtmlEmailChannel
	SMSChannel       *channels.SMSChannel
	PushChannel      *channels.PushChannel
	JobQueue         *queue.JobQueue
	Router           *gin.Engine
}

type NotificationRequest struct {
	Channel string `json:"channel"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Retries int    `json:"retries"`
	Delay   int    `json:"delay"`
}

func NewServer(emailChannel *channels.EmailChannel, htmlEmailChannel *channels.HtmlEmailChannel, smsChannel *channels.SMSChannel, pushChannel *channels.PushChannel, jobQueue *queue.JobQueue) *Server {
	server := &Server{
		EmailChannel:     emailChannel,
		HtmlEmailChannel: htmlEmailChannel,
		SMSChannel:       smsChannel,
		PushChannel:      pushChannel,
		JobQueue:         jobQueue,
		Router:           gin.Default(),
	}
	server.Router.POST("/notify", server.handleNotification)
	return server
}

func (server *Server) handleNotification(ctx *gin.Context) {
	var req NotificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to decode request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var channel channels.NotificationChannel
	switch req.Channel {
	case "email":
		channel = server.EmailChannel
	case "html_email":
		channel = server.HtmlEmailChannel
	case "sms":
		channel = server.SMSChannel
	case "push":
		channel = server.PushChannel
	default:
		log.Error().Msg("Invalid channel: " + req.Channel)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid channel"})
		return
	}

	job := queue.Job{
		Channel: channel,
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
		Retries: req.Retries,
		Delay:   time.Duration(req.Delay) * time.Second,
	}
	server.JobQueue.AddJob(job)
	log.Info().Msg("Job added: " + req.Channel + " to " + req.To)

	ctx.JSON(http.StatusAccepted, gin.H{"message": req.Channel + " notification received and queued successfully"})
}
