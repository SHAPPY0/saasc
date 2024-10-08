package config

import (
	"fmt"
	"os"
	"errors"
	"reflect"
	"path/filepath"
	"github.com/BurntSushi/toml"
	// "github.com/shappy0/saasc/internal/utils"
)

type Conf struct {
	AppName				string		`toml:"AppName"`
	Version				string		`toml:"Version"`
	Commit				string		`toml:"Commit"`
	HomeDir				string		`toml:"Home_Dir"`
	AppDir				string		`toml:"App_Dir"`
	LogLevel			string		`toml:"Log_Level"`
	LogDir				string		`toml:"Log_Dir"`
	LogFilePath			string		`toml:"Log_File_Path"`
	LogDirPath			string		
	ConfigPath			string		
	RefreshRate 		int			`toml:"Refresh_Rate"`
	AzureSubscriptionId	string		`toml:"Azure_Subscription_Id"`
	AzureResourceGroup	string		`toml:"Azure_Resource_Group"`
	AzureClientId		string		`toml:"Azure_Client_Id"`
	AzureClientSecret	string		`toml:"Azure_Client_Secret"`
	AzureTenantId		string		`toml:"Azure_Tenant_Id"`
}

var envVars = map[string]string {
	"AZURE_SUBSCRIPTION_ID":	"AzureSubscriptionId",
	"AZURE_CLIENT_ID":			"AzureClientId",
	"AZURE_CLIENT_SECRET":		"AzureClientSecret",
	"AZURE_TENANT_ID":			"AzureTenantId",	
}

func NewConfig() *Conf {
	homeDir, _ := os.UserHomeDir()
	c := Conf{
		AppName:		AppName,
		HomeDir:		homeDir,
		AppDir:			DefaultAppDir,
		LogLevel:		Info,
		LogDir:			DefaultLogDir,
		LogFilePath:	DefaultFilePath,
		ConfigPath:		filepath.Join(homeDir, DefaultAppDir, DefaultConfigFile),
	}
	return &c
}

func (c *Conf) Load() (*Conf, error) {
	_conf := &Conf{
		AppName:	AppName,
	}
	confPath := c.ConfigPath
	if _, err := os.Stat(confPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("Invalid config file path")
		} else {
			return nil, fmt.Errorf("Unable to load config file")
		}
	}
	_, err := toml.DecodeFile(confPath, &_conf)
	if err != nil {
		return nil, err
	}
	c.SetConfig(_conf)
	setEnvVariabled(_conf)
	return _conf, nil
}

func (c *Conf) SetConfig(conf *Conf) {
	if (conf.HomeDir == "") {
		conf.HomeDir = c.HomeDir
	}
	if (conf.AppDir == "") {
		conf.AppDir = c.AppDir
	}
	if (conf.LogLevel == "") {
		conf.LogLevel = c.LogLevel
	}
	if (conf.LogDir == "") {
		conf.LogDir = c.LogDir
	}
	if (conf.LogFilePath == "") {
		conf.LogFilePath = c.LogFilePath
	}
	if (conf.RefreshRate == 0) {
		conf.RefreshRate = c.RefreshRate
	}
	if (conf.AzureSubscriptionId == "") {
		conf.AzureSubscriptionId = c.AzureSubscriptionId
	}
	if (conf.AzureClientId == "") {
		conf.AzureClientId = c.AzureClientId
	}
	if (conf.AzureClientSecret == "") {
		conf.AzureClientSecret = c.AzureClientSecret
	}
	if (conf.AzureTenantId == "") {
		conf.AzureTenantId = c.AzureTenantId
	}
	if (conf.AzureResourceGroup == "") {
		conf.AzureResourceGroup = c.AzureResourceGroup
	}
	conf.LogDirPath = filepath.Join(conf.HomeDir, conf.AppDir, conf.LogDir, conf.LogFilePath)
}

func (c *Conf) SetResourceGroup(rg string) {
	c.AzureResourceGroup = rg
}

func (c *Conf) GetResourceGroup() string {
	return c.AzureResourceGroup
}

func getValues(config *Conf, key string) string {
	r := reflect.ValueOf(config)
	f := reflect.Indirect(r).FieldByName(key)
	return string(f.String())
}

func setEnvVariabled(config *Conf) {
	for key, value := range envVars {
		os.Setenv(key, getValues(config, value))
	}
}