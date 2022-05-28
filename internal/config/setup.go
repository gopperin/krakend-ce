package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const cmdRoot = "core"

// Setup 载入配置文件
func Setup(path ...string) {

	viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(cmdRoot)
	viper.AddConfigPath(path[0])

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", cmdRoot, err))
		os.Exit(1)
	}

	if err := viper.Unmarshal(&Case); err != nil {
		fmt.Println(err)
	}

}
