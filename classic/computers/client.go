package computers

import (
	"github.com/trustero/jamf-api-client-go/classic/client"
	"net/http"
)

const domain = "computers"

type Service struct {
	client *client.Client
}

func NewService(baseUrl string, username string, password string, httpClient *http.Client) (*Service, error) {

	j, err := client.NewDomainClient(baseUrl, domain, username, password, httpClient)
	if err != nil {
		return nil, err
	}

	return &Service{client: j}, nil
}
