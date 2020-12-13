package init

import "github.com/ggdream/gopack/tools/platform"

type ConfigTemplate struct {
	Name    string `json:"name" yaml:"name" toml:"name"`
	Version string `json:"version" yaml:"version" toml:"version"`

	Author      interface{} `json:"author" yaml:"author" toml:"author"`
	Description string      `json:"description" yaml:"description" toml:"description"`
	Repository  string      `json:"repository" yaml:"repository" toml:"repository"`

	Env    string   `json:"env" yaml:"env" toml:"env"`
	Target []string `json:"target" yaml:"target" toml:"target"`

	Scripts  map[string]string `json:"scripts" yaml:"scripts" toml:"scripts"`
	Packages map[string]string `json:"packages" yaml:"packages" toml:"packages"`
}

func GetConfigStruct(name, version, author, description, repository string) *ConfigTemplate {
	return &ConfigTemplate{
		Name:        name,
		Version:     version,
		Author:      author,
		Description: description,
		Repository:  repository,
		Env:         platform.GetInfo(),
		Target:      []string{platform.GetOSAndArch()},
		Scripts:     map[string]string{},
		Packages:    map[string]string{},
	}
}
