//go:build !((linux || windows) && amd64) && !(linux && (arm || arm64)) && !(windows && arm) && !(darwin && arm64)

package prebuilt

var sqinnName string = "sqinn"
var gzipData []byte = nil
