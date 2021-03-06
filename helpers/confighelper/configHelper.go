package confighelper

import (
	"fmt"

	"github.com/spf13/viper"
)

/*configHelper : Author - Mahesh Chinvar
Purpose : Config file read
*/

//InitViper function to initialize viper
func InitViper() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()
	// Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file %s ", err))
	}
}

//GetConfig method to get configs from config file
func GetConfig(keyName string) string {
	keyValue := viper.GetString(keyName)
	return keyValue
}
