package pack

import (
	"fmt"
	Init "github.com/ggdream/gopack/cmd/init"
	"github.com/ggdream/gopack/global"
	"github.com/ggdream/gopack/tools/caller"
	pack "github.com/ggdream/gopack/tools/compress"
	"github.com/ggdream/gopack/tools/platform"
	"github.com/ggdream/gopack/tools/selects"
	"os"
	"path"
	"strings"
)

func Packs(_type, _path string) error {

	config, err := Init.Parse(_type, _path)
	if err != nil {
		return err
	}

	result := make([]string, 0)
	targets := config.Target
	if targets == nil { // 配置文件没有定义Target字段
		platforms, err := platform.GetPlatform()
		if err != nil {
			return err
		}
		result, err = selects.Select(platforms)
		if err != nil {
			return err
		}
	} else {
		result = targets
	} // 配置文件定义了Target字段

	distPath := path.Join(_path, "dist")
	if _, err := os.Stat(distPath); !os.IsNotExist(err) {
		if err := os.RemoveAll(distPath); err != nil {
			return err
		}
	}

	global.LogMaster.Debug("PACK: the packing work to start\n")
	for _, v := range result {
		global.LogMaster.Info("PACK: pack >> (%s)\n", v)
		if err := Pack(v, config.Name, config.Version); err != nil {
			global.LogMaster.Error("PACK: %s (target: %s)\n", err.Error(), v)
		}
		global.LogMaster.Info("PACK: pack << (%s)\n", v)
	}
	global.LogMaster.Debug("PACK: the packing work to end\n")
	return nil
}

func Pack(target, name, version string) error {
	info := strings.Split(target, "/")
	if err := os.Setenv("GOOS", info[0]); err != nil {
		return err
	}
	if err := os.Setenv("GOARCH", info[1]); err != nil {

		return err
	}

	execPath := fmt.Sprintf("dist/_temp/%s_%s/%s%s",
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
	if msg, err := caller.CallCmdOut("go", "build", "-ldflags", "-s -w", "-o", execPath); err != nil {
		fmt.Printf("target: %s | err: %s\n", target, msg)
		return err
	}

	var packer pack.Packer
	switch info[0] {
	case platform.Windows:
		packer = &pack.ZipPacker{
			DirName: execPath,
			DstName: fmt.Sprintf("dist/%s-%s-%s-%s.zip", name, version, info[0], info[1]),
		}
	default:
		packer = &pack.TgzPacker{
			DirName: execPath,
			DstName: fmt.Sprintf("dist/%s-%s-%s-%s.tgz", name, version, func() string {
				switch info[0] {
				case platform.Darwin:
					return "macos"
				default:
					return info[0]
				}
			}(), info[1]),
		}
	}
	if err := packer.Pack(); err != nil {
		fmt.Printf("compress: %s | err: %s\n", target, err.Error())
		return err
	}

	return nil
}
