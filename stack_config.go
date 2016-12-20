package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var ConfigFile *config

type config struct {
	Image   string
	Command []string
}

//type config struct {
//	Version  int
//	Services map[string]Service
//}
//
//type Service struct {
//	Image string
//}

func initializeStackConfig() {
	viper.SetConfigName("ctest")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Please create a ctest.yaml file: %s \n", err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&ConfigFile)

	if err != nil {
		panic(fmt.Errorf("Unable to parse compose file, %v", err))
	}
}
