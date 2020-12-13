package platform

import (
	"fmt"
	"runtime"
	"strings"
)

func GetOSAndArch() string {
	return strings.Join([]string{runtime.GOOS, runtime.GOARCH}, "/")
}

func GetInfo() string {
	v := runtime.Version()
	_os := runtime.GOOS
	_arch := runtime.GOARCH

	return fmt.Sprintf("%s-%s/%s", v, _os, _arch)
}
