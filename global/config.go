package global

import "github.com/ggdream/gopack/tools/log"

var (
	LogMaster = log.New()

	ConfigFileType = "yaml"
	DefaultVersion = "0.0.1"
	DefaultDesc    = "my new project."
)
