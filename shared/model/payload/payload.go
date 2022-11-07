package payload

import (
	"demo-loginapp2/shared/driver"
)

type Payload struct {
	Data      any                    `json:"data"`
	Publisher driver.ApplicationData `json:"publisher"`
	TraceID   string                 `json:"traceId"`
}
