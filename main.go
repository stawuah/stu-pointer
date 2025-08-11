package main

import (
	"fmt"
	"time"

	"github.com/stawuah/container/container"
	"github.com/stawuah/container/container_resgistry"
	"github.com/stawuah/container/image"
)

// eventHandler is a function that will be registered with the registry.
// It simply prints the event type and the container's ID.
func eventHandler(eventType string, con *container.Container) {
	fmt.Printf("Event received: %s for container ID %s\n", eventType, con.ID)
}

func main() {
	// Initialize a container with some initial values
	myContainer := container.Container{
		Status:    "created",
		IsRunning: false,
		CPU:       1.0,
		Memory:    1024, // 1024 MB
	}

	// Initialize a DockerImage with some starting values
	// The Labels map is initialized to avoid a panic when adding new labels
	myDockerImage := image.DockerImage{
		Repository:   "alpine",
		Tag:          "3.19",
		Size:         100000, // a small size
		LayerCount:   3,
		CreatedAt:    time.Now().Add(-48 * time.Hour), // created 2 days ago
		IsOfficial:   false,
		Downloads:    100,
		Architecture: "amd64",
		Labels:       make(map[string]string),
	}

	fmt.Println("--- Demonstrating Container Methods ---")
	// Print the state before any updates
	fmt.Printf("Before update: Status=%s, IsRunning=%v\n", myContainer.Status, myContainer.IsRunning)

	// Update the status
	myContainer.UpdateStatus("stopped")

	// Print the state after the status update
	fmt.Printf("After update: Status=%s, IsRunning=%v, StartTime=%s\n", myContainer.Status, myContainer.IsRunning, myContainer.StartTime.Format("2006-01-02..."))

	// Add an environment variable
	myContainer.AddEnviromentVars("ENV_VAR", "value")

	// Print the state after adding an environment variable
	fmt.Printf("After env var: Environment map has %d items\n", len(myContainer.Environment))

	// Scale resources
	myContainer.ScaleResources(2.0, 2.0)

	// Print the state after scaling resources
	fmt.Printf("After scaling: CPU=%.1f, Memory=%d\n", myContainer.CPU, myContainer.Memory)

	fmt.Println("\n--- Demonstrating DockerImage Methods ---")

	// --- 1. Call all value receiver methods and print results ---
	fmt.Println("--- State of Image BEFORE calling pointer methods ---")
	fmt.Printf("Full Name: %s\n", myDockerImage.GetFullName())
	fmt.Printf("Size: %d bytes\n", myDockerImage.GetSize())
	fmt.Printf("Is large? %v\n", myDockerImage.IsLarge())
	fmt.Printf("Created At: %s\n", myDockerImage.GetCreatedAt().Format(time.RFC3339))
	fmt.Printf("Is Official? %v\n", myDockerImage.IsOfficialImage())
	fmt.Printf("Age: %s\n", myDockerImage.GetAge().Round(time.Hour))
	fmt.Printf("Has 'maintainer' label? %v\n", myDockerImage.HasLabel("maintainer"))
	fmt.Printf("Current Downloads: %d\n", myDockerImage.Downloads)

	// --- 2. Call all pointer receiver methods ---
	fmt.Println("\n--- Calling pointer receiver methods ---")
	myDockerImage.IncrementDownloads()
	fmt.Println("IncrementDownloads() called.")
	myDockerImage.AddLabel("maintainer", "Alice")
	fmt.Println("AddLabel('maintainer', 'Alice') called.")
	myDockerImage.MarkAsOfficial()
	fmt.Println("MarkAsOfficial() called.")
	myDockerImage.UpdateSize()
	fmt.Println("UpdateSize() called.")

	// --- 3. Call value receiver methods again to see changes ---
	fmt.Println("\n--- State of Image AFTER calling pointer methods ---")
	fmt.Printf("Full Name: %s\n", myDockerImage.GetFullName())
	fmt.Printf("New Size: %d bytes\n", myDockerImage.GetSize())
	fmt.Printf("Is large? %v\n", myDockerImage.IsLarge())
	fmt.Printf("Is Official? %v\n", myDockerImage.IsOfficialImage())
	fmt.Printf("Has 'maintainer' label? %v (value: %s)\n", myDockerImage.HasLabel("maintainer"), myDockerImage.Labels["maintainer"])
	fmt.Printf("New Downloads: %d\n", myDockerImage.Downloads)
	fmt.Printf("Create new image name %v\n", myDockerImage.CreateImageName("stawuah", "1.24.5"))
	fmt.Printf("Full Name after crfeating a new one: %s\n", myDockerImage.GetFullName())

	// Create registry using the constructor
	fmt.Println("\n--- Creating registry ---")
	reg := container_resgistry.NewContainerResgistry(5)

	// Add an event handler that prints events
	reg.AddEventHandler(eventHandler)

	// Create 3 different ContainerConfig instances
	fmt.Println("\n--- Creating container configs ---")
	config1 := &container_resgistry.ContainerConfig{
		Name:  "web-server-1",
		Image: "nginx:latest",
		CPU:   1.0,
	}
	config2 := &container_resgistry.ContainerConfig{
		Name:  "database-server",
		Image: "postgres:13",
		CPU:   2.0,
	}
	config3 := &container_resgistry.ContainerConfig{
		Name:  "web-server-2",
		Image: "nginx:latest",
		CPU:   1.0,
	}

	// Create containers from configs
	fmt.Println("\n--- Creating containers from configs ---")
	c1, err := reg.CreateContainer(config1)
	if err != nil {
		fmt.Printf("Error creating container 1: %v\n", err)
	}
	c2, err := reg.CreateContainer(config2)
	if err != nil {
		fmt.Printf("Error creating container 2: %v\n", err)
	}
	c3, err := reg.CreateContainer(config3)
	if err != nil {
		fmt.Printf("Error creating container 3: %v\n", err)
	}

	// Start 2 containers, stop 1
	fmt.Println("\n--- Starting and stopping containers ---")
	if c1 != nil {
		reg.StartContainer(c1.ID)
	}
	if c2 != nil {
		reg.StartContainer(c2.ID)
	}
	if c3 != nil {
		reg.StopContainer(c3.ID)
	}

	// For demonstration, manually update the status of c1 to "stopped"
	// since you don't have a StopContainer method yet.
	if c1 != nil {
		c1.UpdateStatus("stopped")
		// reg.StopContainer(c3.ID)
	}

	// List running containers
	fmt.Println("\n--- Listing running containers ---")
	runningContainers := reg.ListRunningContainers()
	fmt.Printf("Found %d running container(s):\n", len(runningContainers))
	for _, c := range runningContainers {
		fmt.Printf(" - ID: %s, Name: %s\n", c.ID, c.Name)
	}

	// Print final stats
	fmt.Println("\n--- Final registry stats ---")
	stats := reg.GetStats()
	fmt.Printf("Total containers: %d\n", reg.TotalCount)
	fmt.Printf("Running containers: %d\n", reg.RunningCount)
	fmt.Printf("Registry stats map: %v\n", len(stats))
}
