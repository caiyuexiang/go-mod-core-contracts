//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
)

// IntervalActionResponse defines the Response Content for GET IntervalAction DTOs.
// This object and its properties correspond to the IntervalActionResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/IntervalActionResponse
type IntervalActionResponse struct {
	common.BaseResponse `json:",inline"`
	Action              dtos.IntervalAction `json:"action"`
}

func NewIntervalActionResponse(requestId string, message string, statusCode int, action dtos.IntervalAction) IntervalActionResponse {
	return IntervalActionResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Action:       action,
	}
}

// MultiIntervalActionsResponse defines the Response Content for GET multiple IntervalAction DTOs.
// This object and its properties correspond to the MultiIntervalActionsResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/MultiIntervalActionsResponse
type MultiIntervalActionsResponse struct {
	common.BaseResponse `json:",inline"`
	Actions             []dtos.IntervalAction `json:"actions"`
}

func NewMultiIntervalActionsResponse(requestId string, message string, statusCode int, actions []dtos.IntervalAction) MultiIntervalActionsResponse {
	return MultiIntervalActionsResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Actions:      actions,
	}
}
