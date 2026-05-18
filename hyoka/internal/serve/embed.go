package serve

import (
	"io/fs"

	siteembed "github.com/ronniegeraghty/hyoka/site"
)

// embeddedSite is the served-from root (the contents of site/dist/).
var embeddedSite fs.FS = mustSub(siteembed.FS, "dist")

func mustSub(f fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(f, dir)
	if err != nil {
		panic(err)
	}
	return sub
}
