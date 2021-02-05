package config

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

//SetupViper read the yaml given and parse into values it content
func SetupViper(pathYaml string) {

	viperSetup := viper.GetViper()
	viperSetup.SetConfigType("yml")
	viperSetup.SetConfigName("setup")
	viperSetup.AllowEmptyEnv(false)
	viperSetup.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viperSetup.AutomaticEnv()

	f, err := os.Open(pathYaml)

	if err != nil {
		log.Fatalf("it could not read the yaml file on this path: %s due error: %v", pathYaml, err)
	}

	if err := viper.ReadConfig(bufio.NewReader(f)); err != nil {
		log.Fatalf("it could not parse file: %s due error: %v", pathYaml, err)
	}

	Viper = viperSetup

}
