package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type App struct {
	Server
	Database
}

var AppSetting = &App{}

type Server struct {
	RunMode      string        `yaml:"RunMode"`
	HttpPort     int           `yaml:"HttpPort"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

type Database struct {
	Type        string `yaml:"Type"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	Name        string `yaml:"Name"`
	TablePrefix string `yaml:"TablePrefix"`
}

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("config/app.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, AppSetting)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	AppSetting.Server.ReadTimeout = AppSetting.Server.ReadTimeout * time.Second
	AppSetting.Server.WriteTimeout = AppSetting.Server.WriteTimeout * time.Second
}
