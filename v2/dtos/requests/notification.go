//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// AddNotificationRequest defines the Request Content for POST Notification DTO.
// This object and its properties correspond to the AddNotificationRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-notifications/2.x#/AddNotificationRequest
type AddNotificationRequest struct {
	common.BaseRequest `json:",inline"`
	Notification       dtos.Notification `json:"notification"`
}

// Validate satisfies the Validator interface
func (request AddNotificationRequest) Validate() error {
	err := v2.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddNotificationRequest type
func (request *AddNotificationRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		common.BaseRequest
		Notification dtos.Notification
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*request = AddNotificationRequest(alias)

	// validate AddNotificationRequest DTO
	if err := request.Validate(); err != nil {
		return err
	}
	return nil
}

// AddNotificationReqToNotificationModels transforms the AddNotificationRequest DTO array to the AddNotificationRequest model array
func AddNotificationReqToNotificationModels(reqs []AddNotificationRequest) (n []models.Notification) {
	for _, req := range reqs {
		d := dtos.ToNotificationModel(req.Notification)
		n = append(n, d)
	}
	return n
}

func NewAddNotificationRequest(dto dtos.Notification) AddNotificationRequest {
	return AddNotificationRequest{
		BaseRequest:  common.NewBaseRequest(),
		Notification: dto,
	}
}
