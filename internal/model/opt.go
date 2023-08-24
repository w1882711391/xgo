package model

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
