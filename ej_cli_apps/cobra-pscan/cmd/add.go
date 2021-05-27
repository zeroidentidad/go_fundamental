package cmd

import (
	"fmt"
	"io"
	"os"

	"pscan/scan"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd representa el comando add
var addCmd = &cobra.Command{
	Use:          "add <host1>...<hostn>",
	Aliases:      []string{"a"},
	Short:        "Agregar nuevos hosts a la lista",
	SilenceUsage: true,
	Args:         cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile := viper.GetString("hosts-file")

		return addAction(os.Stdout, hostsFile, args)
	},
}

func addAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostsList{}

	if err := hl.Load(hostsFile); err != nil {
		return err
	}

	for _, h := range args {
		if err := hl.Add(h); err != nil {
			return err
		}

		fmt.Fprintln(out, "Added host:", h)
	}

	return hl.Save(hostsFile)
}

func init() {
	hostsCmd.AddCommand(addCmd)

	// Aquí definir flags y ajustes de configuración.

	// Cobra admite flags persistentes que funcionarán para
	// el comando y todos los subcomandos, por ejemplo:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra admite flags locales que solo se ejecutarán cuando
	// se llame directamente al comando, por ejemplo:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
