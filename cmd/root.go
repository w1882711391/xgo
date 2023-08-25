package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"xgo/internal/model"
)

var rootCmd = &cobra.Command{
	Use:   "xgo",
	Short: "xgo can help you build project Catalog",
	Long:  "xgo can help you build project Catalog,You can read https://github.com/w1882711391/xgo.git",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal("cmd execute err: ", err)
	}
}

func init() {
	model.Init()

	rootCmd.AddCommand(initCmd())
}
