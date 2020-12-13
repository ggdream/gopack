package main

import (
	"fmt"
	"github.com/ggdream/gopack/cli"
	"testing"
)

func TestName(t *testing.T) {
	if err := cli.New(); err != nil {
		fmt.Println(err.Error())
	}
}
