package model

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var Project = &Arch{
	Mutex: &sync.Mutex{},
	Pwd:   "",
	Name:  "",
	Opt: &Opt{
		Redis: Database{Cname: "Redis", Name: "redis"},
		Mysql: Database{Cname: "Mysql", Name: "mysql"},
		Etcd:  Database{Cname: "Etcd", Name: "etcd"},
	},
	Model: make(map[string]*Model),
}

func Init() *Arch {
	var err error

	if err = Project.Read(); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logrus.Fatal("init project err:", err)
		}
		Project.Default()
	}

	return Project
}
