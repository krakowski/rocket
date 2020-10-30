package rocket

const (
	postMessagePath = "api/v1/chat.postMessage"
)

func (message *MessageService) Post(payload MessagePayload) error {
	_, err := message.resty.R().
		SetBody(payload).
		Post(postMessagePath)

	return err
}