package git

import (
	"github.com/ggdream/gopack/tools/caller"
	"strings"
)

func GetUserInfo() (string, string, error) {
	name, err := caller.CallCmdOut("git", "config", "--global", "user.name")
	if err != nil {
		return "", "", err
	}
	email, err := caller.CallCmdOut("git", "config", "--global", "user.email")
	if err != nil {
		return "", "", err
	}

	return purify(name), purify(email), nil
}

func purify(text string) string {
	return strings.Replace(text, "\n", "", -1)
}
