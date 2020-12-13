package run

import (
	"errors"
	"fmt"
	Init "github.com/ggdream/gopack/cmd/init"
	"github.com/ggdream/gopack/global"
	"github.com/ggdream/gopack/tools/caller"
	"strings"
)

func Runner(_type, _path string, script string) error {
	config, err := Init.Parse(_type, _path)
	if err != nil {
		return err
	}

	if code, ok := config.Scripts[script]; ok {
		c := strings.Split(code, " ")
		global.LogMaster.Info("gopack run %s\n", script)
		return caller.CallCmd(c[0], c[1:]...)
	}
	return errors.New(fmt.Sprintf("The script `%s` is not defined", script))
}
