package keycodes

import (
	"github.com/spf13/cobra"
	"prophetstor.com/alameda/datahub/tools/license-utils/pkg/keycodes"
)

var GenerateRegistrationDataCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate registration data",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		keycodes.GenerateRegistrationData()
	},
}

func init() {
}
