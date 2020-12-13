package file

import (
	"errors"
	"os"
)

var (
	yaml = "yaml"
	json = "json"
	toml = "toml"

	app = "gopack"

	types = []string{yaml, json, toml}
)

func JudgeFile() (string, error) {
	for _, v := range types {
		if _, err := os.Stat(app + "." + v); !os.IsNotExist(err) {
			return v, nil
		}
	}
	return "", errors.New("没有找到配置文件")
}
