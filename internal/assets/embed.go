package assets

import (
	"embed"
	"io/fs"
	"log"
)

// Embed the whole /assets directory recursively
//go:embed all:assets
var assetsFS embed.FS

// Sub-filesystems with different roots for use by other packages
// Exposed to outside by functions CowFS(), DocsFS() etc
var (
	cowsFS fs.FS
	docsFS fs.FS
)

// init() runs when package is imported
func init() {
	var err error
	// Make a subFS from /assets/cows
	cowsFS, err = fs.Sub(assetsFS, "assets/cows")
	if err != nil {
		log.Fatalf("Could not create cows subFS: %v", err)
	}
	// Make a subFS from /assets/docs
	docsFS, err = fs.Sub(assetsFS, "assets/docs")
	if err != nil {
		log.Fatalf("Could not create docs subFS: %v", err)
	}
}

func CowsFS() fs.FS {
	return cowsFS
}

func DocsFS() fs.FS {
	return docsFS
}
