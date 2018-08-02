package configuration

import (
	"fmt"
	viper2 "github.com/spf13/viper"
	"log"
	"os"
)

var Env = viper2.New()

func Load() {

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	Env.SetConfigFile(fmt.Sprintf(`%s/authion.toml`, dir))
	err = Env.ReadInConfig()

	if err != nil {
		panic(err)
	}
}
