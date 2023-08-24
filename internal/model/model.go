package model

type Model struct {
	A    *Arch  `json:"-" mapstructure:"-"`
	Name string `json:"name" mapstructure:"name"`
}
