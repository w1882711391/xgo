package model

import (
	"os"
	"path/filepath"
	"xgo/internal/template"
)

func (a *Arch) UtilInit() error {
	var (
		err    error
		ignore string
	)

	if err = os.MkdirAll(filepath.Join(a.Pwd, "util"), 0755); err != nil {
		return err
	}

	if ignore, err = template.Execute("go", "response", nil); err != nil {
		return err
	}

	if err = os.WriteFile(filepath.Join(a.Pwd, "util", "response.go"), []byte(ignore), 0666); err != nil {
		return err
	}

	return nil
}
