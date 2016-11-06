package app

import (
	"os"
	"path/filepath"

	"io/ioutil"

	"strings"

	"github.com/murlokswarm/log"
)

// ResourcePath represents the path of the app resource directory.
type ResourcePath string

// Path is a convenient method that cast a ResourcePath into a string.
func (r ResourcePath) Path() string {
	return string(r)
}

// Join joins any number of path elements into a single path, adding a
// Separator if necessary.
// It calls the Join function from package path/filepath with the resource
// location as the first element.
func (r ResourcePath) Join(elems ...string) string {
	elems = append([]string{r.Path()}, elems...)
	return filepath.Join(elems...)
}

// CSS returns a slice containing the css filenames of the css directory.
func (r ResourcePath) CSS() (css []string) {
	cssPath := r.Join("css")

	info, err := os.Stat(cssPath)
	if err != nil {
		log.Warnf("%v doesn't exists", cssPath)
		return
	}

	if !info.IsDir() {
		log.Errorf("%v is not a directory", cssPath)
		return
	}

	files, _ := ioutil.ReadDir(cssPath)

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		if strings.HasSuffix(f.Name(), ".css") {
			css = append(css, f.Name())
		}
	}

	return
}

// Resources returns the path of the app resource directory.
func Resources() ResourcePath {
	return driver.Resources()
}