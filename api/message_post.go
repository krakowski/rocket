package rocket

import (
	"fmt"
)

const (
	postMessagePath = "api/v1/chat.postMessage"
)

func (message *MessageService) Post(payload MessagePayload) (*MessageResponse, error) {

	var msgError MessageError
	var response MessageResponse
	_, err := message.resty.R().
		SetBody(payload).
		SetError(&msgError).
		SetResult(&response).
		Post(postMessagePath)

	if !response.Success {
		return nil, fmt.Errorf("Posting message to channel %s failed (%s)", payload.Channel, msgError.Error)
	}

	return &response, err
}