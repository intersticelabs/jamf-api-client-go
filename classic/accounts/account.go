// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package accounts

import (
	"context"
	"net/http"

	"github.com/trustero/jamf-api-client-go/classic/client"

	"github.com/pkg/errors"
)

// Accounts returns all system accounts - users and groups
func (j *Service) List() (accounts *JamfAccountsId, response *http.Response, err error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", j.client.Endpoint, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error building JAMF computer query request")
	}

	res := &JamfAccountsResp{}
	if response, err = client.MakeAPIrequest(j.client, req, &res); err != nil {
		err = errors.Wrapf(err, "unable to query accounts from %s", j.client.Endpoint)
		return
	}
	accounts = &res.AccountsId
	return
}

// GetByUserId returns the name, id for a specific user given its Id
func (j *Service) GetByUserId(identifier int) (user *JamfUser, response *http.Response, err error) {
	ep := j.client.UserEndpoint(identifier)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		err = errors.Wrapf(err, "error building JAMF account user request for computer: %v (%s)", identifier, ep)
		return
	}

	res := &JamfUserResp{}
	if response, err = client.MakeAPIrequest(j.client, req, &res); err != nil {
		err = errors.Wrapf(err, "unable to query user for userid : %v (%s)", identifier, ep)
		return
	}
	user = &res.User
	return
}

// GetByGroupId returns the name, id for a specific group given its Id
func (j *Service) GetByGroupId(identifier int) (group *JamfGroup, response *http.Response, err error) {
	ep := j.client.GroupEndpoint(identifier)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		err = errors.Wrapf(err, "error building JAMF account group request for groupid : %v (%s)", identifier, ep)
		return
	}

	res := &JamfGroupResp{}
	if response, err = client.MakeAPIrequest(j.client, req, &res); err != nil {
		err = errors.Wrapf(err, "unable to query account group for groupid: %v (%s)", identifier, ep)
		return
	}
	group = &res.Group
	return
}
