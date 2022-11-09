package models

import "time"

// Metrics is the global data model for the API's.
type Metrics struct {
	MachineID    int       `json:"machineId"`
	Stats        Stats     `json:"stats"`
	LastLoggedIn string    `json:"lastLoggedIn"`
	SysTime      time.Time `json:"sysTime"`
}

// Stats is the child for Metrics which will hold the stats of system.
type Stats struct {
	CPUTemp  int `json:"cpuTemp"`
	FanSpeed int `json:"fanSpeed"`
	HDDSpace int `json:"HDDSpace"`
}
