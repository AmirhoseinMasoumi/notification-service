package channels

type NotificationChannel interface {
	Send(to string, subject string, body string) error
}
