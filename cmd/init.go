package cmd

import (
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"xgo/internal/model"
)

func initCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "init project",
		Long:  "init project , you can use 'wgo init [name]' Build Project Catalog",
		RunE: func(cmd *cobra.Command, args []string) error {
			logrus.Debug("init called")

			if len(args) > 0 && args[0] != "" {
				model.Project.Name = args[0]
			}

			if err := model.Project.Init(); err != nil {
				return err
			}

			color.Green("小汪提示你 初始化项目 %s 成功", model.Project.Name)
			return nil
		},
	}
	return cmd
}
