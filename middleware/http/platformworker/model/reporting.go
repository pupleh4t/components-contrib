package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type SubmitProcessRequest struct {
	DaprAppID   string      `json:"daprAppID"`
	JobUUID     uuid.UUID   `json:"jobUUID"`
	TraceID     string      `json:"traceID"`
	IsSuccess   bool        `json:"isSuccess"`
	ProcessedAt time.Time   `json:"processedAt"`
	Data        interface{} `json:"data"`
}

type SubmitProcessResponse struct {
	Code    int
	Message string
}
