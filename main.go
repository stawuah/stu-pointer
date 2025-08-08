package main

import (
	"fmt"
	"time"

	"github.com/stawuah/container/container"
	"github.com/stawuah/container/image"
)

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
}
