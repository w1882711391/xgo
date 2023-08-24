package model

import (
	"encoding/json"
	"golang.org/x/mod/modfile"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type Arch struct {
	*sync.Mutex

	Name  string            `json:"name" mapstucture:"name"`
	Pwd   string            `json:"pwd" mapstucture:"pwd"`
	Opt   *Opt              `json:"opt" mapstucture:"opt"`
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
	bs, err := os.ReadFile(filepath.Join(pwd, "go.mod"))

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
