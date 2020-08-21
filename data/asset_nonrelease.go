// +build !release

package data

import (
	"net/http"
	"os"
)

// Assets contains project assets.
var Assets http.FileSystem

func init() {
	dir := os.Getenv("TF_DATA")
	if dir == "" {
		dir = "data"
	}
	Assets = http.Dir(dir)
}
