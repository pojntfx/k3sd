package main

import (
	constants "github.com/pojntfx/k3sd/cmd"
	k3sd "github.com/pojntfx/k3sd/pkg/proto/generated"
	"github.com/pojntfx/k3sd/pkg/svc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

const (
	keyPrefix         = "k3sdd."
	configFileDefault = ""
	configFileKey     = keyPrefix + "configFile"
	listenHostPortKey = keyPrefix + "listenHostPort"
)

var rootCmd = &cobra.Command{
	Use:   "k3sd",
	Short: "k3sd is the k3s management daemon",
	Long: `k3sd is the k3s management daemon.

Find more information at:
https://pojntfx.github.io/k3sd/`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("k3sd")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(configFileKey) == configFileDefault) {
			viper.SetConfigFile(viper.GetString(configFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}
		binaryDir := filepath.Join(os.TempDir(), "k3sd")

		listener, err := net.Listen("tcp", viper.GetString(listenHostPortKey))
		if err != nil {
			return err
		}

		server := grpc.NewServer()
		reflection.Register(server)

		extractService := svc.ExtractService{
			BinaryDir:          binaryDir,
			BinaryInternalPath: "/k3s",
		}

		K3SAgentService := svc.K3SAgentManager{
			ExtractService: extractService,
		}
		K3SServerService := svc.K3SServerManager{
			ExtractService: extractService,
		}

		k3sd.RegisterK3SAgentManagerServer(server, &K3SAgentService)
		k3sd.RegisterK3SServerManagerServer(server, &K3SServerService)

		interrupt := make(chan os.Signal, 2)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-interrupt

			// Allow manually killing the process
			go func() {
				<-interrupt

				os.Exit(1)
			}()

			log.Info("Gracefully stopping server (this might take a few seconds)")

			agentMsg := "Could not stop k3s agent"
			serverMsg := "Could not stop k3s server"

			if K3SAgentService.K3SManaged != nil {
				if err := K3SAgentService.K3SManaged.DisableAutoRestart(); err != nil { // Manually disable auto restart; disables crash recovery even if process is not running
					log.Fatal(agentMsg, rz.Err(err))
				}

				if K3SAgentService.K3SManaged.IsRunning() {
					if err := K3SAgentService.K3SManaged.Stop(); err != nil { // Stop is sync, so no need to `.Wait()`
						log.Fatal(agentMsg, rz.Err(err))
					}
				}
			}

			if K3SServerService.K3SManaged != nil {
				if err := K3SServerService.K3SManaged.DisableAutoRestart(); err != nil { // Manually disable auto restart; disables crash recovery even if process is not running
					log.Fatal(serverMsg, rz.Err(err))
				}

				if K3SServerService.K3SManaged.IsRunning() {
					if err := K3SServerService.K3SManaged.Stop(); err != nil { // Stop is sync, so no need to `.Wait()`
						log.Fatal(serverMsg, rz.Err(err))
					}
				}
			}

			if err := extractService.Cleanup(); err != nil {
				log.Fatal(agentMsg, rz.Err(err))
			}

			server.GracefulStop()
		}()

		if err := extractService.Extract(); err != nil {
			return err
		}

		log.Info("Starting server")

		if err := server.Serve(listener); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	var (
		configFileFlag string
		hostPortFlag   string
	)

	rootCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, "Configuration file to use.")
	rootCmd.PersistentFlags().StringVarP(&hostPortFlag, listenHostPortKey, "l", constants.K3SDHostPortDefault, "TCP listen host:port.")

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(constants.CouldNotStartRootCommandErrorMessage, rz.Err(err))
	}
}
