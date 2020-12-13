package platform

import (
	"fmt"
	"testing"
)

func TestGetPlatform(t *testing.T) {
	list, err := GetPlatform()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(list)
}

func TestGetOSAndArch(t *testing.T) {
	println(GetOSAndArch())
}

func TestGetInfo(t *testing.T) {
	println(GetInfo())
}
