package model

import (
	"os"
	"path/filepath"
	"xgo/internal/template"
)

func (a *Arch) EtcInit() error {
	var (
		err  error
		data string
	)

	if err = os.MkdirAll(filepath.Join(a.Pwd, "etc"), 0755); err != nil {
		return err
	}
	if data, err = template.Execute("go", "config.json", nil); err != nil {
		return err
	}

	if err = os.WriteFile(filepath.Join(a.Pwd, "etc", "config.json"), []byte(data), 0666); err != nil {
		return err
	}

	return nil
}
