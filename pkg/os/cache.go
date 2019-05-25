package os

import (
	"github.com/blowaxd/thundersalt/pkg/thunderstore"
	"os"
	"path"
)

const (
	downloadDirectory = "downloads"
	pluginsDirectory  = "extracted-plugins"
	cliDirectory      = ".thundersalt"
)

// GetPluginsCachePath gets the plugins local storage cache
func GetPluginsCachePath() (pluginsCachePath string, err error) {
	pluginsCachePath, err = GetCachePath()
	if err != nil {
		return "", err
	}

	return path.Join(pluginsCachePath, pluginsDirectory), nil
}

// GetPluginDownloadPath Gets the download path where the plugin version will be stored to
func GetPluginDownloadPath(version thunderstore.PluginVersion) (string, error) {
	dir, err := GetPluginsCachePath()
	if err != nil {
		return "", nil
	}
	return path.Join(dir, downloadDirectory, version.FullName, version.ID, ".zip"), nil
}

// GetCachePath gets the thundersalt storage path
func GetCachePath() (cachePath string, err error) {
	cachePath, err = os.UserHomeDir()
	// avoid nesting
	if err == nil {
		return path.Join(cachePath, cliDirectory), nil
	}

	cachePath, err = os.UserCacheDir()
	if err != nil {
		return "", err
	}

	return path.Join(cachePath, cliDirectory), nil
}