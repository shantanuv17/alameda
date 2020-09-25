package app

import (
	"github.com/spf13/cobra"
	"prophetstor.com/alameda/cmd/app"
)

var (
	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display the datahub license-utils version",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			app.PrintSoftwareVer()
		},
	}
)
