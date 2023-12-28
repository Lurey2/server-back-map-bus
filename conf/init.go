package conf

import (
	"fmt"
	"github.com/spf13/viper"
)


func InitializeConfig(){
	
	Migrate()
	initializeViper()
}

func initializeViper(){

	viper.SetConfigName("settings")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}