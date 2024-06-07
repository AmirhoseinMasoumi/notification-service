# Notification Service

This is a simple notification service written in Go (Golang) that allows you to send notifications via email, HTML email, SMS, and push notifications. It provides an HTTP API endpoint for receiving notification requests and queues them for processing.

## Features

- Supports multiple notification channels: email, HTML email, SMS, and push notifications.
- Configurable retry mechanism for failed notifications.
- Asynchronous processing using a job queue.

## Usage

### Sending Notifications

To send a notification, make a POST request to the `/notify` endpoint with the following JSON payload:

```json
{
    "channel": "email",
    "to": "recipient@example.com",
    "subject": "Test Email",
    "body": "This is a test email.",
    "retries": 3,
    "delay": 5
}
```
Replace the "channel" value with the desired notification channel ("email", "html_email", "sms", or "push"), and fill in the "to", "subject", and "body" fields accordingly. You can also specify the number of "retries" and the "delay" in seconds between each retry.

## Configuration
The configuration of the notification service can be adjusted using environment variables. The following variables are supported:

- **ENVIRONMENT**: Set the environment mode (development or production).
- **SMTP_HOST: SMTP**: server host address.
- **SMTP_PORT: SMTP**: server port number.
- **SMTP_USERNAME**: SMTP username.
- **SMTP_PASSWORD**: SMTP password.
- **SMS_API_KEY**: API key for SMS service provider.
- **PUSH_API_KEY**: API key for push notification service provider.
- **JOB_QUEUE_SIZE**: Size of the job queue.
- **JOB_WORKER_COUNT**: Number of worker threads for processing jobs.
- **SERVER_ADDRESS**: Address and port on which the server should listen.

### Usage Notes
For SMS and push notifications, you need to implement the respective logic in the channels/SMS.go and channels/Push.go files. Replace the placeholder implementation with the actual logic for sending SMS and push notifications.

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/notification-service.git
    cd notification-service
    ```
2. Set up environment variables by creating a .env file:
    ```sh
    cp .env.example .env
    ```
    Edit the .env file with your configuration settings.

3. Build and run the application:
    ```sh
    go build -o notification-service main.go
    ./notification-service
    ```
## Contributing
Contributions are welcome! Please fork the repository and create a pull request. For major changes, open an issue to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (git checkout -b feature/YourFeature)
3. Commit your changes (git commit -m 'Add some feature')
4. Push to the branch (git push origin feature/YourFeature)
5. Open a pull request

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
