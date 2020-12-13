package caller

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func CallCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if len(args) > 2 && args[1] == "init" {
		sl := strings.Split(args[2], "/")
		cmd.Dir = sl[len(sl)-1]
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return cmd.Wait()
}

func CallCmdOut(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var dst bytes.Buffer
	cmd.Stdout = &dst

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return dst.String(), nil
}
