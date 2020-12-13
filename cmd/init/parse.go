package init

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ggdream/gopack/global"
	"github.com/naoina/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type Parser struct {
	Data []byte
}

func (p *Parser) FromJson() (*ConfigTemplate, error) {
	var config ConfigTemplate
	err := json.Unmarshal(p.Data, &config)

	if err != nil {
		return nil, err
	}
	return &config, err
}

func (p *Parser) FromYaml() (*ConfigTemplate, error) {
	var config ConfigTemplate
	err := yaml.Unmarshal(p.Data, &config)

	if err != nil {
		return nil, err
	}
	return &config, err
}

func (p *Parser) FromToml() (*ConfigTemplate, error) {
	var config ConfigTemplate
	err := toml.Unmarshal(p.Data, &config)

	if err != nil {
		return nil, err
	}
	return &config, err
}

func Parse(_type, _path string) (*ConfigTemplate, error) {
	file, err := os.Open(path.Join(_path, fmt.Sprintf("%s.%s", global.NAME, _type)))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	parser := Parser{Data: data}
	switch _type {
	case YAML:
		return parser.FromYaml()
	case JSON:
		return parser.FromJson()
	case TOML:
		return parser.FromToml()
	default:
		return nil, errors.New("the file format error")
	}
}
