package keycodes

import (
	"github.com/spf13/cobra"
	"prophetstor.com/alameda/datahub/tools/license-utils/pkg/keycodes"
)

var ReadKeycodeCmd = &cobra.Command{
	Use:   "read",
	Short: "read keycode detailed information",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		keycodes.ListKeycodes(keycode)
	},
}

func init() {
	parseReadKeycodeFlag()
}

func parseReadKeycodeFlag() {
	ReadKeycodeCmd.Flags().StringVar(&keycode, "keycode", "", "Read the specified keycode information")
}
