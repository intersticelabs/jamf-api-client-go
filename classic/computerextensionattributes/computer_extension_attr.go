// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package computerextensionattributes

import (
	"context"
	"github.com/trustero/jamf-api-client-go/classic/client"
	"github.com/pkg/errors"
	"net/http"
)

// ComputerExtensionAttrExists is a helper function to check if an extension attribute
// exists without having to parse the response. Note: If an error occurs that doesn't include
// a not found message ... we log the error and return false
//func (j *Service) ComputerExtensionAttrExists(identifier interface{}) bool {
//	_, err := j.ComputerExtensionAttributeDetails(identifier)
//	if err != nil {
//		if !strings.Contains(err.Error(), "The server has not found anything matching the request URI") {
//			log.Error().Msgf("did not find computer extension attribute %v due to %s", identifier, err.Error())
//		}
//		return false
//	}
//	return true
//}

// ComputerExtensionAttributes returns all computer extension attributes
func (j *Service) ComputerExtensionAttributes() ([]ComputerExtensionAttribute, error) {

	req, err := http.NewRequestWithContext(context.Background(), "GET", j.client.Endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building JAMF computer extension attribute query request")
	}

	res := &ComputerExtensionAttributes{}
	if _, err := client.MakeAPIrequest(j.client, req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query computer extension attribute from %s", j.client.Endpoint)
	}
	return res.List, nil
}

// ComputerExtensionAttributeDetails returns the details for a specific computer extension attribute given its Id or Name
//func (j *Service) ComputerExtensionAttributeDetails(identifier interface{}) (*ComputerExtensionAttributeDetails, error) {
//	ep, err := EndpointBuilder(j.Endpoint, computerExtAttrContext, identifier)
//	if err != nil {
//		return nil, errors.Wrapf(err, "error building JAMF query request endpoint for script: %v", identifier)
//	}
//	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
//	if err != nil {
//		return nil, errors.Wrapf(err, "error building JAMF query request for computer extension attribute: %v", identifier)
//	}
//
//	res := ComputerExtensionAttributeDetails{}
//	if err := j.makeAPIrequest(req, &res); err != nil {
//		return nil, errors.Wrapf(err, "unable to query computer extension attribute with Id: %v from %s", identifier, ep)
//	}
//
//	return &res, nil
//}

// UpdateComputerExtensionAttribue will update a computer extension attribute in Jamf by either Id or Name
//func (j *Service) UpdateComputerExtensionAttribue(identifier interface{}, content *ComputerExtensionAttribute) (*ComputerExtensionAttribute, error) {
//	ep, err := EndpointBuilder(j.Endpoint, computerExtAttrContext, identifier)
//	if err != nil {
//		return nil, errors.Wrapf(err, "error building JAMF query request for computer extension attribute: %v", identifier)
//	}
//
//	err = ValidateComputerExtensionAttribute(content)
//	if err != nil {
//		return nil, errors.Wrapf(err, "computer extension attribute validation failed: %v", identifier)
//	}
//
//	bodyContent, err := xml.Marshal(content)
//	if err != nil {
//		return nil, errors.Wrapf(err, "error building JAMF update payload for computer extension attribute: %v", identifier)
//	}
//
//	body := bytes.NewReader(bodyContent)
//	req, err := http.NewRequestWithContext(context.Background(), "PUT", ep, body)
//	if err != nil {
//		return nil, errors.Wrapf(err, "error building JAMF update request for computer extension attribute: %v (%s)", identifier, ep)
//	}
//
//	res := ComputerExtensionAttribute{}
//	if err := j.makeAPIrequest(req, &res); err != nil {
//		return nil, errors.Wrapf(err, "unable to process JAMF update request for computer extension attribute: %v (%s)", identifier, ep)
//	}
//
//	return &res, nil
//}

// CreateComputerExtensionAttribute will create a computer extension attribute in Jamf
//func (j *Service) CreateComputerExtensionAttribute(content *ComputerExtensionAttribute) (*ComputerExtensionAttribute, error) {
//	// -1 denotes the next available Id
//	ep, err := EndpointBuilder(j.Endpoint, computerExtAttrContext, -1)
//	if err != nil {
//		return nil, errors.Wrapf(err, "error building JAMF query request for new computer extension attribute")
//	}
//
//	if content == nil {
//		return nil, errors.Wrapf(fmt.Errorf("Empty payload"), "unable to process JAMF creation request for computer extension attribute: (%s)", ep)
//	}
//
//	if content.Name == "" {
//		return nil, errors.Wrapf(fmt.Errorf("Name required for new computer extension attribute"), "unable to process JAMF creation request for computer extension attribute: (%s)", ep)
//	}
//
//	err = ValidateComputerExtensionAttribute(content)
//	if err != nil {
//		return nil, errors.Wrapf(err, "computer extension attribute validation failed: %v", content.Name)
//	}
//
//	bodyContent, err := xml.Marshal(content)
//	if err != nil {
//		return nil, errors.Wrapf(err, "error building JAMF creation payload for computer extension attribute: %v", content.Name)
//	}
//
//	body := bytes.NewReader(bodyContent)
//	req, err := http.NewRequestWithContext(context.Background(), "POST", ep, body)
//	if err != nil {
//		return nil, errors.Wrapf(err, "error building JAMF creation request for computer extension attribute: %v (%s)", content.Name, ep)
//	}
//
//	res := ComputerExtensionAttribute{}
//	if err := j.makeAPIrequest(req, &res); err != nil {
//		return nil, errors.Wrapf(err, "unable to process JAMF creation request for computer extension attribute: %v (%s)", content.Name, ep)
//	}
//
//	return &res, nil
//}

// DeleteComputerExtensionAttribute will delete a computer extension attribute by either Id or Name
//func (j *Service) DeleteComputerExtensionAttribute(identifier interface{}) (*ComputerExtensionAttribute, error) {
//	ep, err := EndpointBuilder(j.Endpoint, computerExtAttrContext, identifier)
//	if err != nil {
//		return nil, errors.Wrapf(err, "error building JAMF query request for computer extension attribute: %v", identifier)
//	}
//
//	req, err := http.NewRequestWithContext(context.Background(), "DELETE", ep, nil)
//	if err != nil {
//		return nil, errors.Wrapf(err, "error building JAMF deletion request for computer extension attribute: %v (%s)", identifier, ep)
//	}
//
//	res := ComputerExtensionAttribute{}
//	if err := j.makeAPIrequest(req, &res); err != nil {
//		return nil, errors.Wrapf(err, "unable to process JAMF deletion request for computer extension attribute: %v (%s)", identifier, ep)
//	}
//
//	return &res, nil
//}
