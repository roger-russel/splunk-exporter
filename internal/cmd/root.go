package cmd

import (
	"fmt"

	"github.com/roger-russel/splunk-exporter/internal/cmd/dto"
	"github.com/roger-russel/splunk-exporter/internal/config"
	"github.com/roger-russel/splunk-exporter/internal/splunk"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

//Root command executer
func Root(vf dto.FullVersion) {

	flags := dto.Flags{}

	rootCmd = &cobra.Command{
		Use:   "splunk-exporter",
		Short: "splunk-exporter",
		Run: func(cmd *cobra.Command, args []string) {
			splunk.GetData(&flags)
		},
	}

	rootCmd.AddCommand(Version(vf))

	rootCmd.Flags().StringVarP(&flags.Config, "config", "c", "sexporter.yaml", "the config yaml: -c sexporter.yaml")
	rootCmd.Flags().StringVarP(&flags.User, "user", "u", "", "user that will be used to login on splunk")
	rootCmd.Flags().StringVarP(&flags.Password, "password", "p", "", "password that will be used to login on splunk")

	config.SetupViper(flags.Config)

	err := rootCmd.Execute()

	if err != nil {
		fmt.Println("fail to execute due error", err)
	}

}
