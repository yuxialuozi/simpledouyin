package biz

import httpStatus "github.com/cloudwego/hertz/pkg/protocol/consts"

var (
	InvalidActionType = GWError{HTTPStatusCode: httpStatus.StatusBadRequest, StatusCode: 400008, StatusMsg: "invalid action type"}
)
