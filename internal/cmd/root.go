package cmd

import (
	"fmt"
	"os"

	v "github.com/roger-russel/splunk-exporter/internal/cmd/version"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

//Root command executer
func Root(vf v.FullVersion) {
	checkDefaultCommand()

	rootCmd = &cobra.Command{
		Use:   "splunk-exporter",
		Short: "splunk-exporter",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.AddCommand(Version(vf))
	err := rootCmd.Execute()

	if err != nil {
		fmt.Println("fail to execute due error", err)
	}

}

func checkDefaultCommand() {
	if len(os.Args) < 2 {
		os.Args = append([]string{os.Args[0], "--help"}, os.Args[1:]...)
	}
}
