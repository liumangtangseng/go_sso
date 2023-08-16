package core

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sso/config"
)

const configFile = "settings.yaml"

// ConfInit 读取yaml配置
func ConfInit() config.Serve {
	c := &config.Serve{}
	yamlConf, err := ioutil.ReadFile(configFile)
	if err != nil {
		//panic(fmt.Errorf("get yamlConf error: %s", err))
		logrus.Fatalf("get yamlConf error: %s", err)
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		logrus.Fatalf("config Init Unmarshal: %v", err)
	}
	logrus.Println("config yamlFile load Init success.")
	return *c
}
