// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package client

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	scriptsContext         = "scripts"
	ComputersContext       = "computers"
	computerExtAttrContext = "computerextensionattributes"
	policiesContext        = "policies"
)

// Client represents the interface used to communicate with
// the Jamf API via an HTTP client
type Client struct {
	Domain   string
	Username string
	Password string
	Endpoint string
	logger   *logrus.Logger
	Api      *http.Client
}

// Used if custom client not passed on when NewDomainClient instantiated
func DefaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Minute,
	}
}

// NewDomainClient returns a new Jamf HTTP client to be used for API requests
func NewDomainClient(baseUrl string, domain string, username string, password string, client *http.Client) (*Client, error) {
	if baseUrl == "" || username == "" || password == "" {
		return nil, errors.New("you must provide a valid Jamf base url, username, and password")
	}

	if client == nil {
		client = DefaultHTTPClient()
	}

	return &Client{
		Domain:   baseUrl,
		Username: username,
		Password: password,
		Endpoint: fmt.Sprintf("%s/JSSResource/%s", baseUrl, domain),
		Api:      client,
	}, nil
}

func (j *Client) makeAPIrequest(r *http.Request, v interface{}) (*http.Response, error) {
	return MakeAPIrequest(j, r, v)
}

// MockAPIRequest is used for testing the API client
func (j *Client) MockAPIRequest(r *http.Request, v interface{}) (*http.Request, error) {
	r.Header.Set("Accept", "application/json,  application/xml;q=0.9")
	r.SetBasicAuth(j.Username, j.Password)
	_, err := j.makeAPIrequest(r, v)
	return r, err
}

// EndpointBuilder can be utilized to query a specific API context via name
func (j *Client) NameEndpoint(identifier string) string {
	return fmt.Sprintf("%s/name/%s", j.Endpoint, identifier)
}

// EndpointBuilder can be utilized to query a specific API context via Id
func (j *Client) IdEndpoint(identifier int) string {
	return fmt.Sprintf("%s/id/%d", j.Endpoint, identifier)
}

func MakeAPIrequest(j *Client, r *http.Request, v interface{}) (*http.Response, error) {
	// Jamf API only sends XML for some endpoints so we will accept both but prioritize
	// JSON responses with the quallity value of 1.0 and 0.9 for XML responses
	// https://developer.mozilla.org/en-US/docs/Glossary/quality_values
	r.Header.Set("Accept", "application/json, application/xml;q=0.9")
	r.Header.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0, post-check=0, pre-check=0")
	r.Header.Set("Strict-Transport-Security", "max-age=31536000 ; includeSubDomains")
	r.SetBasicAuth(j.Username, j.Password)

	res, err := j.Api.Do(r)
	if err != nil {
		return res, errors.Wrapf(err, "error making %s request to %s", r.Method, r.URL)
	}
	defer res.Body.Close()

	// If status code is not ok attempt to read the response in plain text
	if res.StatusCode != 200 && res.StatusCode != 201 {
		var responseData []byte
		responseData, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return res, errors.Wrapf(err, "request error: %s. unable to retrieve plain text response: %s", res.Status, err.Error())
		}
		return res, fmt.Errorf("request error: %s", string(responseData))
	}

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options
	// ex. [text/xml charset=UTF-8]
	contentType := strings.Split(res.Header.Get("Content-Type"), ";")
	switch t := contentType[0]; t {
	case "text/xml", "application/xml":
		if err = xml.NewDecoder(res.Body).Decode(&v); err != nil {
			// TODO: return a string or something
			return res, errors.Wrapf(err, "response was successful but error occured decoding response body of type %s", t)
		}
	case "text/json", "application/json", "text/plain":
		if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
			return res, errors.Wrapf(err, "response was successful but error occured error decoding response body of type %s", t)
		}
	default:
		return res, errors.Wrapf(err, "response was successful but error occured recieved unexpected response body of type %s", t)
	}

	return res, nil
}
