package model

import (
	"os"
	"path/filepath"
	"xgo/internal/template"
)

func (a *Arch) daoInit() error {
	var (
		ignore string
		err    error
	)

	if err = os.MkdirAll(filepath.Join(a.Pwd, "dao"), 0755); err != nil {
		return err
	}

	if ignore, err = template.Execute("go", "mysql", nil); err != nil {
		return err
	}

	if err = os.WriteFile(filepath.Join(a.Pwd, "dao", "mysql.go"), []byte(ignore), 0666); err != nil {
		return err
	}

	if ignore, err = template.Execute("go", "redis", nil); err != nil {
		return err
	}

	if err = os.WriteFile(filepath.Join(a.Pwd, "dao", "redis.go"), []byte(ignore), 0666); err != nil {
		return err
	}
	return nil
}
