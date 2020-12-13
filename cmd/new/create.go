package new

import (
	"fmt"
	Init "github.com/ggdream/gopack/cmd/init"
	"github.com/ggdream/gopack/tools/caller"
	"os"
	"path"
	"strings"
)

func getModuleSuffix(module string) string {
	names := strings.Split(module, "/")
	return names[len(names)-1]
}

func Create(_type, _path, module string, name, version, author, description, repository string) error {
	filePath := path.Join(_path, getModuleSuffix(module))
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		var entry string
		fmt.Printf("The directory `%s` already exists, do you want to overwrite it? (y/n):", getModuleSuffix(module))

		for {
			_, err := fmt.Scanln(&entry)
			if err != nil {
				fmt.Print("Please entry `y` or `n` :")
				continue
			}

			switch entry {
			case "y":
				if err := os.RemoveAll(filePath); err != nil {
					return err
				}
				goto Label
			case "n":
				os.Exit(0)
			default:
				fmt.Print("Please entry `y` or `n` :")
			}
		}
	}

	// create it!
Label:
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}

	if err := caller.CallCmd("go", "mod", "init", repository); err != nil {
		return err
	}
	return Init.Generate(_type, filePath, name, version, author, description, repository)
}
