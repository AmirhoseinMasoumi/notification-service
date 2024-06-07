package queue

import (
	"time"

	"github.com/AmirhoseinMasoumi/notification-service/internal/channels"
)

type Job struct {
	Channel channels.NotificationChannel
	To      string
	Subject string
	Body    string
	Retries int
	Delay   time.Duration
}
