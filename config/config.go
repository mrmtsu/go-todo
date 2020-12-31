package config

import (
	"log"

	"gopkg.in/ini.v1"
)

// ConfigList ...
type ConfigList struct {
	Port      string
	SQLDirver string
	DBName    string
	LogFile   string
}

// Config ...
var Config ConfigList

func init() {
	LoadConfig()
}

// LoadConfig ...
func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8000"),
		SQLDirver: cfg.Section("db").Key("driver").String(),
		DBName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
