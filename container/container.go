package container

import (
	"errors"
	"time"
)

type Container struct {
	ID          string
	Name        string
	Image       string
	Status      string
	Port        int
	CPU         float64
	Memory      int64
	IsRunning   bool
	StartTime   time.Time
	Environment map[string]string
}

func (con *Container) UpdateStatus(new_status string) {
	if new_status == "" { // Check for an empty string
		errors.New("pass a a status or no status passed")
		return // Add a return to stop execution
	}

	con.Status = new_status // Assign the new status to the field

	switch new_status {
	case "running": // Corrected spelling
		con.IsRunning = true
		con.StartTime = time.Now()
	case "stopped":
		con.IsRunning = false
	}
}

func (con *Container) AddEnviromentVars(key, value string) {

	if con.Environment == nil {
		con.Environment = make(map[string]string)

	}

	con.Environment[key] = value

}

func (con *Container) ScaleResources(cpuMultiplier float64, memoryMultiplier float64) {

	con.CPU *= cpuMultiplier
	con.Memory *= int64(memoryMultiplier)

}
