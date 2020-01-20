package cmd

import (
	constants "github.com/pojntfx/go-isc-dhcp/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "k3sserverctl",
	Short: "k3sserverctl manages servers on k3sd, the k3s management daemon",
	Long: `k3sserverctl manages servers on k3sd, the k3s management daemon.

Find more information at:
https://pojntfx.github.io/k3sd/`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("k3sserver")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
}

// Execute starts the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(constants.CouldNotStartRootCommandErrorMessage, rz.Err(err))
	}
}
