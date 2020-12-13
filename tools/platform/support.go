package platform

import (
	"github.com/ggdream/gopack/tools/caller"
	"strings"
)

const (
	Windows = "windows"
	Darwin  = "darwin"
	Linux   = "linux"
)

func GetPlatform() ([]string, error) {
	res, err := caller.CallCmdOut("go", "tool", "dist", "list")
	if err != nil {
		return nil, err
	}

	data := strings.Split(res, "\n")

	return filter(data[:len(data)-1]), nil
}

func filter(data []string) (final []string) {
	for _, v := range data {
		if strings.HasPrefix(v, Windows) || strings.HasPrefix(v, Linux) || strings.HasPrefix(v, Darwin) {
			final = append(final, v)
		}
	}
	return
}
