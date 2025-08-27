//go:build darwin && arm64

package prebuilt

import (
	_ "embed"
)

var sqinnName string = "sqinn"

//go:embed "darwin-arm64.gz"
var gzipData []byte
