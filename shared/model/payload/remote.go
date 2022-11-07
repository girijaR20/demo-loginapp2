package payload

import "demo-loginapp2/shared/driver"

type Args struct {
	Type      string                 `json:"type"`
	Data      any                    `json:"data"`
	Publisher driver.ApplicationData `json:"publisher"`
	TraceID   string                 `json:"traceId"`
}

type Reply struct {
	Success      bool
	ErrorMessage string
	Data         any
}
