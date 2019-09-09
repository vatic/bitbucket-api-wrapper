package client

import (
	"fmt"

	"github.com/spf13/viper"
)

const cfgFile string = "config.json"

func ReposList(projectName string) {
	initConfig(cfgFile)
}

func initConfig(cfgFile string) {
	viper.SetConfigFile(cfgFile)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	fmt.Println(viper.Get("url"))
}
