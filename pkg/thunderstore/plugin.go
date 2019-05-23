package thunderstore

import (
	"errors"
	"time"
)

type Plugin struct {
	ID         string          `json:"uuid4"`
	Name       string          `json:"name"`
	FullName   string          `json:"full_name"`
	Owner      string          `json:"owner"`
	PackageURL string          `json:"package_url"`
	IsPinned   bool            `json:"is_pinned"`
	Versions   []PluginVersion `json:"versions"`
}

type PluginVersion struct {
	ID            string    `json:"uuid4"`
	Name          string    `json:"name"`
	FullName      string    `json:"full_name"`
	Description   string    `json:"description"`
	IconURL       string    `json:"icon"`
	VersionNumber string    `json:"version_number"`
	Dependencies  []string  `json:"dependencies"`
	DownloadUrl   string    `json:"download_url"`
	Downloads     uint64    `json:"downloads"`
	DateCreated   time.Time `json:"date_created"`
	WebsiteURL    string    `json:"website_url"`
	IsActive      bool      `json:"is_active"`
}

var NoPluginVersionError = errors.New("The given plugin contains no version")

// GetLatestVersion
func GetLatestVersion(plugin Plugin) (PluginVersion, error) {
	if len(plugin.Versions) <= 0 {
		return PluginVersion{}, NoPluginVersionError
	}
	return plugin.Versions[0], nil
}
