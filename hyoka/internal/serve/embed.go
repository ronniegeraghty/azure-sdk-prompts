package serve

import "embed"

//go:embed all:site
var embeddedSite embed.FS
