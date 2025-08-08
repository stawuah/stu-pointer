package image

import (
	"math/rand/v2"
	"time"
)

type DockerImage struct {
	Repository   string
	Tag          string
	Size         int64
	LayerCount   int
	CreatedAt    time.Time
	IsOfficial   bool
	Downloads    int64
	Architecture string
	Labels       map[string]string
	Digest       string
}

type DockerImageInterface interface {
	GetFullName() string
	GetSize() int64
	CreateImageName() (string, string)
	GetCreatedAt() time.Time
	IsOfficialImage() bool
	IncrementDownloads()
	AddLabel(key, value string)
	IsLarge() bool
	GetAge() time.Duration
	HasLabel(key string) bool
	MarkAsOfficial()
	UpdateSize()
}

// GetFullName returns "repository:tag"
func (img DockerImage) GetFullName() string {
	return img.Repository + ":" + img.Tag
}

func (img DockerImage) GetSize() int64 {
	return img.Size
}

// IsLarge returns true if size > 1GB (1073741824 bytes)
func (img DockerImage) IsLarge() bool {

	const maxSize = 1073741824 // 1 GB in bytes
	return img.Size > maxSize
}

// GetCreatedAt returns the creation timestamp of the image.
func (img DockerImage) GetCreatedAt() time.Time {
	return img.CreatedAt
}

// IsOfficialImage returns true if the image is an official build.
func (img DockerImage) IsOfficialImage() bool {
	return img.IsOfficial
}

// GetAge returns time since CreatedAt
func (img DockerImage) GetAge() time.Duration {

	return time.Since(img.CreatedAt)
}

// HasLabel checks if a label exists
func (img DockerImage) HasLabel(key string) bool {
	_, has_label := img.Labels[key]

	return has_label
}

// IncrementDownloads increases download count
func (img *DockerImage) IncrementDownloads() {
	img.Downloads++
}

// AddLabel adds or updates a label
func (img *DockerImage) AddLabel(key, value string) {

	// Initialize the Labels map if it's nil to prevent a panic.
	if img.Labels == nil {
		img.Labels = make(map[string]string)
	}

	img.Labels[key] = value
}

// MarkAsOfficial sets IsOfficial to true and adds "official" label
func (img *DockerImage) MarkAsOfficial() {

	img.IsOfficial = true
	img.AddLabel("vendor", "official")
	// img.Labels[official_label] = official_label

}

// UpdateSize recalculates size (simulate with random value for now)
func (img *DockerImage) UpdateSize() {
	maxSize := int64(1073741824)
	img.Size = rand.Int64N(maxSize)
}

func (img *DockerImage) CreateImageName(repository_name, tag string) (string string) {
	img.Repository = repository_name
	img.Tag = tag
	return img.Repository + ":" + img.Tag
}
