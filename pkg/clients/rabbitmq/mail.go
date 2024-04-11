package rabbitmq


func SendMail( email, message, subject string) {
	msgs := Messages{
		Email:    email,
		Messages: message,
		Subject:  subject,
	}
	PublishNotifications(msgs)
}