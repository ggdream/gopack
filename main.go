package main

import (
	"fmt"
	"github.com/ggdream/gopack/cli"
)

func main() {
	if err := cli.New(); err != nil {
		fmt.Println(err.Error())
	}
}
