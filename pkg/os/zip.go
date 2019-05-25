package os

import (
	"archive/zip"
	"errors"
	"github.com/blowaxd/thundersalt/pkg/thunderstore"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Unzip unzips the given plugin to the cache directory
// Zip will be extracted to : cachePath + plugin.FullName + plugin.VersionID
func Unzip(src string, plugin thunderstore.PluginVersion) ([]string, error) {
	cachePath, err := GetPluginsCachePath()
	if err != nil {
		log.Fatalln("error getting the PluginsCachePath")
		return nil, err
	}
	dest := path.Join(cachePath, plugin.FullName, plugin.ID)

	var filenames []string

	r, err := zip.OpenReader(src)
	defer r.Close()
	if err != nil {
		return filenames, err
	}

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: https://snyk.io/research/zip-slip-vulnerability#go
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, errors.New(fpath + ": illegal file path")
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			_ = os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		err = outFile.Close()
		if err != nil {
			return filenames, err
		}
		err = rc.Close()
		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
