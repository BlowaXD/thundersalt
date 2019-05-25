package os

import (
	"github.com/blowaxd/thundersalt/pkg/thunderstore"
	"io"
	"net/http"
	"os"
)

// DownloadPlugin will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadPlugin(plugin thunderstore.PluginVersion) error {
	// Get the data
	resp, err := http.Get(plugin.DownloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	zipFilePath, err := GetPluginDownloadPath(plugin)
	if err != nil {
		return err
	}
	out, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
