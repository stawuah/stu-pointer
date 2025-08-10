package containerresgistry

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/stawuah/container/container"
)

type ContainerConfig struct {
	Name          string
	Image         string
	Ports         []int
	Environment   map[string]string
	CPU           float64
	Memory        int64
	Volumes       []string
	Networks      []string
	Labels        map[string]string
	RestartPolicy string
}

type ContainerRegistry struct {
	containers    map[string]*container.Container
	totalCount    int
	runningCount  int
	maxContainers int
	createdAt     time.Time
	lastUpdate    time.Time
	stats         map[string]int
	eventHandlers []func(string, *container.Container)
}

func NewContainerResgistry(maxContainers int) *ContainerRegistry {

	return &ContainerRegistry{
		containers:    make(map[string]*container.Container),
		totalCount:    0,
		runningCount:  0,
		maxContainers: 0,
		createdAt:     time.Now(),
		lastUpdate:    time.Now(),
		stats:         make(map[string]int),
		eventHandlers: make([]func(string, *container.Container), 0), // ask wuestion here when you are done that how do you use literals to initailise slice and map
	}

}

// CreateContainer creates container from config and adds to registry
func (r *ContainerRegistry) CreateContainer(config *ContainerConfig) (*container.Container, error) {
	// Check if we've hit maxContainers limit
	if r.totalCount >= r.maxContainers {
		return nil, errors.New("maximum container limit reached")
	}

	// Generate unique ID
	id := fmt.Sprintf("%x-%d", rand.Int63(), time.Now().UnixNano())

	// Create Container from config
	new_container := &container.Container{
		ID:          id,
		Name:        config.Name,
		Image:       config.Image,
		Status:      "created", // Containers will be "created" by default
		CPU:         config.CPU,
		Memory:      config.Memory,
		Environment: config.Environment,
		Port:        0, // Assuming a single port for now.
		IsRunning:   false,
		StartTime:   time.Time{}, // Not running yet, so StartTime is zero value
	}
	// Add to containers map
	r.containers[new_container.ID] = new_container
	// Update totalCount
	r.totalCount++

	r.lastUpdate = time.Now()

	// Call event handlers with "created" event
	for _, handler := range r.eventHandlers {
		handler("created", new_container)
	}
	// Return pointer to created container

	return new_container, nil
}

// StartContainer starts container by ID
func (r *ContainerRegistry) StartContainer(id string) error {
	// Find container by ID (return error if not found)
	connect_contained, ok := r.containers[id]

	if !ok {
		return errors.New("container not found")
	}
	// Update container status to "running"

	connect_contained.Status = "running"
	// Increment runningCount
	r.runningCount++
	// Update lastUpdate time
	r.lastUpdate = time.Now()
	// Call event handlers with "started" event
	for _, running_container := range r.eventHandlers {
		running_container("started", connect_contained)
	}

	return nil
}

func (r *ContainerRegistry) StopContainer(id string) error {
	// Find container by ID (return error if not found)
	connect_contained, ok := r.containers[id]

	if !ok {
		return errors.New("container not found")
	}
	// Update container status to "running"

	connect_contained.Status = "stopped"
	// Increment runningCount
	r.runningCount++
	// Update lastUpdate time
	r.lastUpdate = time.Now()
	// Call event handlers with "started" event
	for _, running_container := range r.eventHandlers {
		running_container("started", connect_contained)
	}

	return nil
}

func (r *ContainerRegistry) GetContainer(id string) (*container.Container, error) {
	// Return pointer to container or error if not found
	found_container, ok := r.containers[id]

	if !ok {
		return nil, errors.New("container not found")
	}

	return found_container, nil

}

// ListRunningContainers returns slice of pointers to running containers
func (r *ContainerRegistry) ListRunningContainers() []*container.Container {
	// Filter containers by status and return slice of pointers
	var runningContainers []*container.Container

	for _, running_container := range r.containers {
		// Check if the container's status is "running".
		if running_container.IsRunning {
			// If it is, append the container pointer to our slice.
			runningContainers = append(runningContainers, running_container)
		}
	}

	return runningContainers

}

func (r *ContainerRegistry) AddEventHandler(handler func(string, *container.Container)) {
	// Add handler to eventHandlers slice
	r.eventHandlers = append(r.eventHandlers, handler)
}

func (r *ContainerRegistry) GetStats() map[string]int {
	// Return copy of stats map (not pointer!)

	statsCopy := make(map[string]int, len(r.stats))

	for key, value := range r.stats {
		statsCopy[key] = value
	}

	return statsCopy
}
