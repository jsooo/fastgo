package config

import (
	"os"
	"strings"

	"github.com/astaxie/beego/config"
)

func Reader(filePath string) (iniconf config.Configer, err error) {
	path, _ := os.Getwd()
	if strings.Contains(path, "service") {
		path += "/../.."
	}
	appconf, err := config.NewConfig("ini", path+"/config/app.conf")
	if err != nil {
		return
	}
	runmode := appconf.String("runmode")
	if runmode == "dev" {
		iniconf, err = config.NewConfig("ini", path+"/config/dev/"+filePath)
	} else {
		iniconf, err = config.NewConfig("ini", path+"/config/prod/"+filePath)
	}
	return iniconf, err
}
