package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/GptPluginHub/hub/cmd/app/options"
	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/version"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var onlyOneSignalHandler = make(chan struct{})

func NewServerCommand() *cobra.Command {
	viper.AutomaticEnv()
	o := options.NewDefaultOption()

	cmd := &cobra.Command{
		Use:  "hub",
		Long: `The hub server is chatgpt plugin hub service.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			defer close(onlyOneSignalHandler) // panics when called twice
			if conf, err := config.TryLoadConfig(o.ConfigFile); err == nil {
				o = &options.Options{
					ServerOptions: o.ServerOptions,
					Conf:          conf,
				}
			}
			ctx, cancel := context.WithCancel(context.Background())
			c := make(chan os.Signal, 2)
			signal.Notify(c, []os.Signal{os.Interrupt, syscall.SIGTERM}...)
			go func() {
				<-c
				cancel()
				<-c
				os.Exit(1) // second signal. Exit directly.
			}()
			server := o.NewServer(ctx)
			return server.Run(ctx)
		},
	}
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of hub server",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(version.Get())
		},
	}
	cmd.AddCommand(versionCmd)

	fs := cmd.Flags()
	fs.AddFlagSet(o.Flags())

	return cmd
}
