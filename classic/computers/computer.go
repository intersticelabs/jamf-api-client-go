// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package computers

import (
	"context"
	"fmt"
	"github.com/intersticelabs/jamf-api-client-go/classic/client"
	"net/http"

	"github.com/pkg/errors"
)

type ListResponse struct {
	Size      int              `json:"size" xml:"size"`
	Computers []ComputerNameId `json:"computers,omitempty" xml:"computers,omitempty"`
}

// Computers returns all enrolled computer devices
func (j *Service) List() (computers []ComputerNameId, response *http.Response, err error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", j.client.Endpoint, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error building JAMF computer query request")
	}

	res := &ListResponse{}
	if response, err = client.MakeAPIrequest(j.client, req, &res); err != nil {
		err = errors.Wrapf(err, "unable to query enrolled computers from %s", j.client.Endpoint)
		return
	}
	computers = res.Computers
	return
}

// Computers returns all enrolled computer devices
func (j *Service) ListWithBasicInfo() (result []BasicComputerInfo, response *http.Response, err error) {
	ep := fmt.Sprintf("%s/subset/basic", j.client.Endpoint)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		err = errors.Wrap(err, "error building JAMF computer query request")
		return
	}

	res := &Computers{}
	if response, err = client.MakeAPIrequest(j.client, req, &res); err != nil {
		err = errors.Wrapf(err, "unable to query enrolled computers from %s", ep)
	}
	result = res.List
	return
}

type getByIdResponse struct {
	Computer Computer `json:"computer"`
}

// GetById returns the details for a specific computer given its Id
func (j *Service) GetById(identifier int) (result *Computer, response *http.Response, err error) {
	ep := j.client.IdEndpoint(identifier)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		err = errors.Wrapf(err, "error building JAMF computer request for computer: %v (%s)", identifier, ep)
		return
	}

	res := &getByIdResponse{}
	if response, err = client.MakeAPIrequest(j.client, req, &res); err != nil {
		err = errors.Wrapf(err, "unable to query enrolled computer for computer: %v (%s)", identifier, ep)
		return
	}
	result = &res.Computer
	return
}

// GetById returns the details for a specific computer given its Id
func (j *Service) GetHardwareByUid(uid string) (result *HardwareInformation, response *http.Response, err error) {
	ep := fmt.Sprintf("%s/serialnumber/%s/subset/Hardware", j.client.Endpoint, uid)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		err = errors.Wrapf(err, "error building JAMF computer request for computer: %v (%s)", uid, ep)
		return
	}

	res := &getByIdResponse{}
	if response, err = client.MakeAPIrequest(j.client, req, &res); err != nil {
		err = errors.Wrapf(err, "unable to query enrolled computer for computer: %v (%s)", uid, ep)
		return
	}
	result = &res.Computer.Hardware
	return
}
