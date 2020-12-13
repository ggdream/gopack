package build

import (
	"fmt"
	Init "github.com/ggdream/gopack/cmd/init"
	"github.com/ggdream/gopack/global"
	"github.com/ggdream/gopack/tools/caller"
	"github.com/ggdream/gopack/tools/platform"
	"os"
	"path"
	"strings"
)

func contain(el string, array []string) bool {
	for _, v := range array {
		if el == v {
			return true
		}
	}
	return false
}

func Builds(_type, _path string) error {
	config, err := Init.Parse(_type, _path)
	if err != nil {
		return err
	}

	targets := config.Target
	if targets == nil {
		return Build(platform.GetOSAndArch(), config.Name)
	}

	platforms, err := platform.GetPlatform()
	if err != nil {
		return err
	}

	distPath := path.Join(_path, "dist")
	if _, err := os.Stat(distPath); !os.IsNotExist(err) {
		if err := os.RemoveAll(distPath); err != nil {
			return err
		}
	}

	global.LogMaster.Debug("BUILD: the compilation work to start\n")
	for _, target := range targets {
		if contain(target, platforms) {
			global.LogMaster.Info("BUILD: compile >> (%s)\n", target)
			if err := Build(target, config.Name); err != nil {
				global.LogMaster.Error("BUILD: %s (target: %s)\n", err.Error(), target)
				continue
			}
			global.LogMaster.Info("BUILD: compile << (%s)\n", target)
		} else {
			global.LogMaster.Error("BUILD: the `%s` is not supported\n", target)
		}
	}
	global.LogMaster.Debug("BUILD: the compilation work to end\n")

	return nil
}

func Build(target, name string) error {
	info := strings.Split(target, "/")
	if err := os.Setenv("GOOS", info[0]); err != nil {
		return err
	}
	if err := os.Setenv("GOARCH", info[1]); err != nil {

		return err
	}

	_path := fmt.Sprintf("dist/_temp/%s_%s/%s%s",
		info[0],
		info[1],
		name,
		func() string {
			switch info[0] {
			case platform.Windows:
				return ".exe"
			default:
				return ""
			}
		}(),
	)
	if _, err := caller.CallCmdOut("go", "build", "-o", _path); err != nil {
		return err
	}

	return nil
}
