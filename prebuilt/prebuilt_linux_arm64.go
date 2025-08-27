//go:build linux && arm64

package prebuilt

import (
	_ "embed"
)

var sqinnName string = "sqinn"

//go:embed "linux-arm64.gz"
var gzipData []byte
