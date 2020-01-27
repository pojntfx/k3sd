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

var getKubeconfigCmd = &cobra.Command{
	Use:     "kubeconfig",
	Aliases: []string{"k"},
	Short:   "Get a kubeconfig",
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

		kubeconfig, err := client.GetKubeconfig(ctx, &k3sd.K3SServerEmptyArgs{})
		if err != nil {
			return err
		}

		fmt.Print(kubeconfig.GetKubeconfig())

		return nil
	},
}

func init() {
	getKubeconfigCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.K3SDHostPortDefault, constants.HostPortDocs)
	getKubeconfigCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, constants.ConfigurationFileDocs)

	if err := viper.BindPFlags(getKubeconfigCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	getCmd.AddCommand(getKubeconfigCmd)
}
