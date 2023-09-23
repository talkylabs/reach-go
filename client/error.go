// Package error provides the interface for Reach specific errors.
package client

import (
	"encoding/json"
	"fmt"
)

// ReachRestError provides information about an unsuccessful request.
type ReachRestError struct {
	Code     int                    `json:"errorCode"`
	Details  map[string]interface{} `json:"errorDetails"`
	Message  string                 `json:"errorMessage"`
	MoreInfo string                 `json:"more_info"`
	Status   int                    `json:"status"`
}

func (err *ReachRestError) Error() string {
	detailsJSON, _ := json.Marshal(err.Details)
	return fmt.Sprintf("Status: %d - ApiError %d: %s (%s) More info: %s",
		err.Status, err.Code, err.Message, detailsJSON, err.MoreInfo)
}
