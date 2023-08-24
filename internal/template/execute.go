package template

import (
	"fmt"
	"html/template"
	"strings"
	"xgo/internal"
	"xgo/internal/opt"
)

// Execute 将目标文件的内容读取出来
func Execute(dir string, name string, data any) (string, error) {
	tplpwd := fmt.Sprintf("template/%s/%s.tpl", dir, name)

	bs, err := internal.Template.ReadFile(tplpwd)

	if err != nil {
		return "", err
	}

	t, err := template.New(name).Funcs(opt.FuncMap()).Parse(string(bs))

	if err != nil {
		return "", err
	}

	write := strings.Builder{}

	if err := t.Execute(&write, data); err != nil {
		return "", err
	}

	return write.String(), nil
}
