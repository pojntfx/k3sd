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
	Short:   "Start a server",
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

		client := k3sd.NewK3SServerManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if _, err := client.Start(ctx, &k3sd.K3SServer{
			NetworkDevice: viper.GetString(networkDeviceKey),
			TLSSan:        viper.GetString(tlsSanKey),
		}); err != nil {
			return err
		}

		fmt.Println("server started")

		return nil
	},
}

func init() {
	var (
		serverHostPortFlag string
		configFileFlag     string
		networkDeviceFlag  string
		tlsSanFlag         string
	)

	startCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.K3SDHostPortDefault, "Host:port of the k3sd server to use.")
	startCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, "Configuration file to use.")
	startCmd.PersistentFlags().StringVarP(&networkDeviceFlag, networkDeviceKey, "d", "edge0", "The name of the network device to use.")
	startCmd.PersistentFlags().StringVarP(&tlsSanFlag, tlsSanKey, "i", "192.168.1.10", "Additional IP to generate TLS certificates for.")

	if err := viper.BindPFlags(startCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(startCmd)
}
