package tidy

import (
	"github.com/ggdream/gopack/tools/caller"
	"testing"
)

func TestCallCmd(t *testing.T) {
	if err := caller.CallCmd("ping", "golang.google.cn"); err != nil {
		panic(err.Error())
	}
}

func TestFetch(t *testing.T) {
	if err := Fetch("yaml", "./"); err != nil {
		panic(err.Error())
	}
}

func TestTidy(t *testing.T) {
	err := Tidy("yaml", "./")
	if err != nil {
		panic(err)
	}
}
