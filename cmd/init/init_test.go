package init

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	err := Generate(YAML, "./", "app", "1.0.0", "mocaraka@gmail.com", "test", "github.com/ggdream/home")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestParse(t *testing.T) {
	config, err := Parse(YAML, "./gopack.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(*config)
}
