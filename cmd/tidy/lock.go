package tidy

import (
	"fmt"
	Init "github.com/ggdream/gopack/cmd/init"
	"github.com/ggdream/gopack/tools/caller"
)

func Fetch(_type, _path string) error {
	config, err := Init.Parse(_type, _path)
	if err != nil {
		return err
	}

	for name, tag := range config.Packages {
		module := fmt.Sprintf("%s@%s", name, tag)
		if err := caller.CallCmd("go", "get", module); err != nil {
			continue
		}
	}
	return nil
}

func Tidy(_type, _path string) error {
	if err := caller.CallCmd("go", "mod", "tidy"); err != nil {
		return err
	}

	return Fetch(_type, _path)
}
