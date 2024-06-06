package templates

import (
	"encoding/json"
	"log"
	"os"
)

type ViteManifest map[string]struct {
	File     string
	CssFiles []string `json:"css"`
}

func ParseManifest(path string) (ViteManifest, error) {
	binary, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var manifest ViteManifest
	err = json.Unmarshal(binary, &manifest)
	return manifest, err
}

var IsProd = os.Getenv("GIN_ENV") == "production"
var Manifest ViteManifest

func init() {
	if !IsProd {
		return
	}

	manifest, err := ParseManifest("./priv/assets/.vite/manifest.json")
	if err != nil {
		log.Fatal(err)
	}
	Manifest = manifest
}
