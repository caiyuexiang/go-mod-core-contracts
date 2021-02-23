//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
)

// IntervalResponse defines the Response Content for GET Interval DTOs.
// This object and its properties correspond to the IntervalResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/IntervalResponse
type IntervalResponse struct {
	common.BaseResponse `json:",inline"`
	Interval            dtos.Interval `json:"interval"`
}

func NewIntervalResponse(requestId string, message string, statusCode int, interval dtos.Interval) IntervalResponse {
	return IntervalResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Interval:     interval,
	}
}

// MultiIntervalsResponse defines the Response Content for GET multiple Interval DTOs.
// This object and its properties correspond to the MultiIntervalsResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/MultiIntervalsResponse
type MultiIntervalsResponse struct {
	common.BaseResponse `json:",inline"`
	Intervals           []dtos.Interval `json:"intervals"`
}

func NewMultiIntervalsResponse(requestId string, message string, statusCode int, intervals []dtos.Interval) MultiIntervalsResponse {
	return MultiIntervalsResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Intervals:    intervals,
	}
}
