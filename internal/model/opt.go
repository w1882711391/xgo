package model

import (
	"os"
	"path/filepath"
	"xgo/internal/template"
)

type Database struct {
	Enable bool   `json:"enable" mapstructure:"enable"`
	Cname  string `json:"cname" mapstructure:"cname"`
	Name   string `json:"name" mapstructure:"name"`
}

type Opt struct {
	Redis Database `json:"redis" mapstructure:"redis"`
	Mysql Database `json:"mysql" mapstructure:"mysql"`
	Etcd  Database `json:"etcd" mapstructure:"etcd"`
}

func (a *Arch) OptInit() error {
	var (
		err    error
		ignore string
	)

	if err = os.MkdirAll(filepath.Join(a.Pwd, "opt"), 0755); err != nil {
		return err
	}

	if ignore, err = template.Execute("go", "opt", nil); err != nil {
		return err
	}

	if err = os.WriteFile(filepath.Join(a.Pwd, "opt", "opt.go"), []byte(ignore), 0666); err != nil {
		return err
	}
	return nil
}
