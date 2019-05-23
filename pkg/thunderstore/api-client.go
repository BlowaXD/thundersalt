package thunderstore

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	apiUrl         = "https://thunderstore.io/api/v1/"
	packageListUrl = apiUrl + "package/"
)

var PluginNotFoundError = errors.New("Plugin was not found")

// GetPlugins Gets all plugins available
func GetPlugins() ([]Plugin, error) {
	resp, err := http.Get(packageListUrl)
	defer resp.Body.Close()
	if err != nil {
		return []Plugin{}, err
	}
	var plugins []Plugin

	err = json.NewDecoder(resp.Body).Decode(&plugins)

	if err != nil {
		return []Plugin{}, err
	}
	return plugins, nil
}

// GetPluginByName Gets the plugin by the given name
func GetPluginByName(name string) (Plugin, error) {
	plugins, err := GetPlugins()
	if err != nil {
		return Plugin{}, err
	}
	for _, plugin := range plugins {
		if plugin.Name == name {
			return plugin, nil
		}
	}
	return Plugin{}, PluginNotFoundError
}

// GetPluginByID Gets a plugin by its ID
func GetPluginByID(ID string) (Plugin, error) {
	plugins, err := GetPlugins()
	if err != nil {
		return Plugin{}, err
	}
	for _, plugin := range plugins {
		if plugin.ID == ID {
			return plugin, nil
		}
	}
	return Plugin{}, PluginNotFoundError
}

//
func GetLatestDownloadByPluginID(ID string) (string, error) {
	plugin, err := GetPluginByID(ID)
	if err != nil {
		return "", err
	}
	// Versions[0] is supposed to be the latest version of the plugin
	return GetLatestVersion(plugin), nil
}
