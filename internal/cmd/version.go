package cmd

import (
	"fmt"

	"github.com/roger-russel/splunk-exporter/internal/cmd/dto"
	"github.com/spf13/cobra"
)

//Version of the binary built
func Version(vf dto.FullVersion) (versionCmd *cobra.Command) {

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version of splunk-exporter",
		Long:  `Print information about this build of splunk-exporter`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version: %s\nbuilded at: %s\ncommit hash: %s\n", vf.Version, vf.Date, vf.Commit)
		},
	}

	return versionCmd
}
