package cmd

import (
	"context"
	"fmt"

	constants "github.com/pojntfx/k3sd/cmd"
	k3sd "github.com/pojntfx/k3sd/pkg/proto/generated"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"s"},
	Short:   "Start an agent",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(configFileKey) == configFileDefault) {
			viper.SetConfigFile(viper.GetString(configFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}

		conn, err := grpc.Dial(viper.GetString(serverHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := k3sd.NewK3SAgentManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if _, err := client.Start(ctx, &k3sd.K3SAgent{
			NetworkDevice: viper.GetString(networkDeviceKey),
			Token:         viper.GetString(tokenKey),
			ServerURL:     viper.GetString(serverUrlKey),
		}); err != nil {
			return err
		}

		fmt.Println("agent started")

		return nil
	},
}

func init() {
	var (
		networkDeviceFlag string
		tokenFlag         string
		serverUrlFlag     string
	)

	startCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.K3SDHostPortDefault, constants.HostPortDocs)
	startCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, constants.ConfigurationFileDocs)
	startCmd.PersistentFlags().StringVarP(&networkDeviceFlag, networkDeviceKey, "d", "edge0", "The name of the network device to use.")
	startCmd.PersistentFlags().StringVarP(&tokenFlag, tokenKey, "t", "asdf", "The token to authenticate with.")
	startCmd.PersistentFlags().StringVarP(&serverUrlFlag, serverUrlKey, "u", "https://:6443", "The URL of the server to connect to.")

	if err := viper.BindPFlags(startCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(startCmd)
}
