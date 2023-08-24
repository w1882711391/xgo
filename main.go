package main

import (
	"github.com/sirupsen/logrus"
	"os/exec"
)

func main() {
	defer func() {
		var err error

		if err = exec.Command("go", "mod", "tidy").Run(); err != nil {
			logrus.Debugf("fina run 'go mod tidy' err: %v", err)
		}

		if err = exec.Command("go", "fmt", "./...").Run(); err != nil {
			logrus.Debugf("fina run 'go fmt ./...' err: %v", err)
		}
	}()
}
