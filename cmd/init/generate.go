package init

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ggdream/gopack/global"
	"github.com/naoina/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Generator struct {
	Config *ConfigTemplate
}

func (g *Generator) ToYaml() ([]byte, error) { return yaml.Marshal(g.Config) }
func (g *Generator) ToToml() ([]byte, error) { return toml.Marshal(g.Config) }

func (g *Generator) ToJson() ([]byte, error) {
	raw, err := json.Marshal(g.Config)
	if err != nil {
		return nil, err
	}

	var data bytes.Buffer
	err = json.Indent(&data, raw, "", "\t")

	return data.Bytes(), err
}

func Generate(_type, _path string, name, version, author, description, repository string) error {
	configStruct := GetConfigStruct(name, version, author, description, repository)
	generator := &Generator{
		Config: configStruct,
	}

	var err error
	data := make([]byte, 0)
	switch _type {
	case YAML:
		data, err = generator.ToYaml()
	case JSON:
		data, err = generator.ToJson()
	case TOML:
		data, err = generator.ToToml()
	default:
		err = errors.New(fmt.Sprintf("The file format `%s` is not supported", _type))
	}
	if err != nil {
		return err
	}

	fileName := strings.Join([]string{global.NAME, _type}, ".")
	filePath := path.Join(_path, fileName)
	if _, err = os.Stat(filePath); !os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("The file `%s` already exist", fileName))
	}

	return ioutil.WriteFile(filePath, data, 0644)
}
