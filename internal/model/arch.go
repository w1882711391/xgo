package model

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/mod/modfile"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"xgo/internal/template"
)

type Arch struct {
	*sync.Mutex

	Name  string            `json:"name" mapstucture:"name"`
	Pwd   string            `json:"pwd" mapstucture:"pwd"`
	Opt   *Opt              `json:"opt.tpl" mapstucture:"opt.tpl"`
	Model map[string]*Model `json:"model" mapstucture:"model"`
}

// Default 给pwd和name赋值
func (a *Arch) Default() {
	var (
		pwd, _ = os.Getwd()
	)

	a.Pwd = pwd
	a.Name = filepath.Base(pwd)

	//读取go.mod文件 为了接下来拿到项目名称
	bs, err := os.ReadFile(path.Join(pwd, "go.mod"))

	if err != nil {
		return
	}
	//解析go.mod文件
	mf, err := modfile.Parse("", bs, nil)
	if err != nil {
		return
	}

	if mod := mf.Module.Mod.Path; mod != "" {
		a.Name = strings.Split(mod, "/")[len(strings.Split(mod, "/"))-1]
	}
}

// Write 将Arch解析成json并写入.wgo.json中
func (a *Arch) Write() error {
	bs, _ := json.MarshalIndent(a, "", "    ")
	return os.WriteFile(path.Join(a.Pwd, ".wgo.json"), []byte(bs), 0755)
}

// Read 将.wgo.json读取解析json文件
func (a *Arch) Read() error {
	var err error

	viper.SetConfigType("json")
	viper.SetConfigFile("./.wgo.json")

	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	if err = viper.Unmarshal(a); err != nil {
		return fmt.Errorf("unmarshal .wgo.json err: %w", err)
	}
	return nil
}

func (a *Arch) Init() error {
	var (
		err    error
		ignore string
	)

	if ignore, err = template.Execute("go", "main", nil); err != nil {
		return err
	}

	if err = os.WriteFile(path.Join(a.Pwd, "main.go"), []byte(ignore), 0666); err != nil {
		return err
	}

	if err = exec.Command("git", "init", "-q").Run(); err != nil {
		return err
	}

	if ignore, err = template.Execute("git", "gitignore", nil); err != nil {
		return err
	}

	if err = os.WriteFile(path.Join(a.Pwd, ".gitignore"), []byte(ignore), 0666); err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Join(a.Pwd, "model"), 0755); err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Join(a.Pwd, "route"), 0755); err != nil {
		return err
	}

	if ignore, err = template.Execute("go", "route", nil); err != nil {
		return err
	}

	if err = os.WriteFile(path.Join(a.Pwd, "route", "route.go"), []byte(ignore), 0666); err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Join(a.Pwd, "service"), 0755); err != nil {
		return err
	}

	if err = a.daoInit(); err != nil {
		return err
	}

	if err = a.OptInit(); err != nil {
		return err
	}

	if err = a.UtilInit(); err != nil {
		return err
	}

	if err = a.EtcInit(); err != nil {
		return err
	}

	if err = exec.Command("go", "mod", "init", a.Name).Run(); err != nil {
		logrus.Debugf("run go mod init %s err: %v", a.Name, err)
	}

	return nil
}
