package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Season-04/appos/core/manifest"
)

func (cfg *Config) Load() error {
	err := os.MkdirAll(filepath.Join(cfg.Directory, "apps.enabled"), 0777)
	if err != nil {
		return err
	}

	files, err := filepath.Glob(filepath.Join(cfg.Directory, "apps.enabled/*.json"))
	if err != nil {
		return err
	}

	cfg.mtx.Lock()
	defer cfg.mtx.Unlock()

	cfg.manifests = make([]manifest.Manifest, 0, len(files))
	cfg.errors = make([]ManifestError, 0)

	for _, file := range files {
		manifest := manifest.Manifest{}
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, &manifest)
		if err != nil {
			cfg.errors = append(cfg.errors, ManifestError{Path: file, Error: err})
		} else {
			cfg.manifests = append(cfg.manifests, manifest)
		}
	}

	return cfg.ensureRequestAppsAreEnabled()
}
