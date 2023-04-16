package service

import (
	"encoding/json"
	"io"
)

type Client struct {
	config            ClientConfig
	requestBuilder    requestBuilder
	createFormBuilder func(writer io.Writer) formBuilder
}

func NewClient(authToken string) *Client {
	config := DefaultConfig(authToken)
	return NewClientWithConfig(config)
}

// NewClientWithConfig creates new OpenAI API client for specified config.
func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config:         config,
		requestBuilder: newRequestBuilder(),
		createFormBuilder: func(body io.Writer) formBuilder {
			return newFormBuilder(body)
		},
	}
}

type marshaller interface {
	marshal(value any) ([]byte, error)
}

type jsonMarshaller struct {
}

func (jm *jsonMarshaller) marshal(value any) ([]byte, error) {
	return json.Marshal(value)
}
