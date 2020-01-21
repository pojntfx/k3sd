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

var cleanupCmd = &cobra.Command{
	Use:     "cleanup",
	Aliases: []string{"c"},
	Short:   "Clean up an agent",
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

		if _, err := client.Cleanup(ctx, &k3sd.K3SAgentEmptyArgs{}); err != nil {
			return err
		}

		fmt.Println("agent cleaned up")

		return nil
	},
}

func init() {
	var (
		serverHostPortFlag string
		configFileFlag     string
	)

	cleanupCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.K3SDHostPortDefault, constants.HostPortDocs)
	cleanupCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, constants.ConfigurationFileDocs)

	if err := viper.BindPFlags(cleanupCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(cleanupCmd)
}
