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
func (j *Service) List() ([]ComputerNameId, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", j.client.Endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building JAMF computer query request")
	}

	res := &ListResponse{}
	if err := client.MakeAPIrequest(j.client, req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query enrolled computers from %s", j.client.Endpoint)
	}
	return res.Computers, nil
}

// Computers returns all enrolled computer devices
func (j *Service) ListWithBasicInfo() ([]BasicComputerInfo, error) {
	ep := fmt.Sprintf("%s/subset/basic", j.client.Endpoint)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building JAMF computer query request")
	}

	res := &Computers{}
	if err := client.MakeAPIrequest(j.client, req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query enrolled computers from %s", ep)
	}
	return res.List, nil
}

type getByIdResponse struct {
	Computer Computer `json:"computer"`
}

// GetById returns the details for a specific computer given its Id
func (j *Service) GetById(identifier int) (*Computer, error) {
	ep := j.client.IdEndpoint(identifier)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF computer request for computer: %v (%s)", identifier, ep)
	}

	res := &getByIdResponse{}
	if err := client.MakeAPIrequest(j.client, req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query enrolled computer for computer: %v (%s)", identifier, ep)
	}
	return &res.Computer, nil
}

// GetById returns the details for a specific computer given its Id
func (j *Service) GetHardwareByUid(uid string) (*HardwareInformation, error) {
	ep := fmt.Sprintf("%s/serialnumber/%s/subset/Hardware", j.client.Endpoint, uid)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF computer request for computer: %v (%s)", uid, ep)
	}

	res := &getByIdResponse{}
	if err := client.MakeAPIrequest(j.client, req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query enrolled computer for computer: %v (%s)", uid, ep)
	}
	return &res.Computer.Hardware, nil
}
