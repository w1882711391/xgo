package opt

import (
	"html/template"
	"strings"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"replace": strings.Replace,
	}
}
