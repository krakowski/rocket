package rocket

import "errors"

const (
	postMessagePath = "api/v1/chat.postMessage"
)

func (message *MessageService) Post(payload MessagePayload) (*MessageResponse, error) {

	var response MessageResponse
	_, err := message.resty.R().
		SetBody(payload).
		SetResult(&response).
		Post(postMessagePath)

	if !response.Success {
		return nil, errors.New("Posting message failed")
	}

	return &response, err
}