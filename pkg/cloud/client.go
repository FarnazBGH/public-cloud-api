package cloud

import (
	"context"

	apiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

type Client struct {
	apiClient *apiclient.APIClient
	ctx       context.Context
}

func NewClient(apiKey, apiHost string) *Client {
	ctx := context.WithValue(
		context.Background(),
		apiclient.ContextAPIKeys,
		map[string]apiclient.APIKey{
			"X-LSW-Auth": {Key: apiKey},
		},
	)

	configuration := apiclient.NewConfiguration()
	configuration.Host = apiHost
	apiClient := apiclient.NewAPIClient(configuration)

	return &Client{
		apiClient: apiClient,
		ctx:       ctx,
	}
}
