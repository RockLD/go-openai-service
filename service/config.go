package service

import (
	"github.com/spf13/viper"
	"net/http"
)

type APIType string

const (
	APITypeOpenAI             APIType = "OPEN_AI"
	APITypeAzure              APIType = "AZURE"
	APITypeAzureAD            APIType = "AZURE_AD"
	defaultEmptyMessagesLimit uint    = 300
)

type ClientConfig struct {
	authToken string

	BaseURL    string
	OrgID      string
	APIType    APIType
	APIVersion string // required when APIType is APITypeAzure or APITypeAzureAD
	Engine     string // required when APIType is APITypeAzure or APITypeAzureAD

	HTTPClient *http.Client

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		authToken:          authToken,
		BaseURL:            viper.GetString("openaiAPIURL"),
		APIType:            APITypeOpenAI,
		OrgID:              "",
		HTTPClient:         &http.Client{},
		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}
