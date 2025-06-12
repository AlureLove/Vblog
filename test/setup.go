package test

import (
	"github.com/infraboard/mcube/v2/ioc"
	"os"
	"path/filepath"
)

func DevelopmentSetup() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = findConfigFile()
	}
	req.ConfigFile.Path = path
	err := ioc.ConfigIocObject(req)
	if err != nil {
		panic(err)
	}
}

func findConfigFile() string {
	wd, err := os.Getwd()
	if err != nil {
		return "etc/application.toml"
	}

	dir := wd
	for {
		configPath := filepath.Join(dir, "etc", "application.toml")
		if _, err := os.Stat(configPath); err == nil {
			relPath, err := filepath.Rel(wd, configPath)
			if err != nil {
				return configPath
			}
			return relPath
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "etc/application.toml"
}
