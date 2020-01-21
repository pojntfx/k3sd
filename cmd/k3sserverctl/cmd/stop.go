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

var stopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"d"},
	Short:   "Stop a server",
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

		if _, err := client.Stop(ctx, &k3sd.K3SServerEmptyArgs{}); err != nil {
			return err
		}

		fmt.Println("server stopped")

		return nil
	},
}

func init() {
	var (
		serverHostPortFlag string
		configFileFlag     string
	)

	stopCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.K3SDHostPortDefault, constants.HostPortDocs)
	stopCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, constants.ConfigurationFileDocs)

	if err := viper.BindPFlags(stopCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(stopCmd)
}
