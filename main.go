package main

import (
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"xgo/cmd"
	"xgo/internal/model"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logrus.SetReportCaller(true)
}
func main() {

	defer func() {
		var err error

		if err = exec.Command("go", "mod", "tidy").Run(); err != nil {
			logrus.Debugf("fina run 'go mod tidy' err: %v", err)
		}

		if err = exec.Command("go", "fmt", "./...").Run(); err != nil {
			logrus.Debugf("fina run 'go fmt ./...' err: %v", err)
		}

		if err = model.Project.Write(); err != nil {
			logrus.Warnf("fina write .wgo.json err: %v", err)
		}
	}()

	if err := exec.Command("go", "version").Run(); err != nil {
		color.Red("command 'go' not found")
		os.Exit(1)
	}

	cmd.Execute()
}
