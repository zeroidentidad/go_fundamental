package cmd

import (
	"fmt"
	"io"
	"os"

	"pscan/scan"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "Listar hosts en la lista de hosts",
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile := viper.GetString("hosts-file")

		return listAction(os.Stdout, hostsFile, args)
	},
}

func listAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostsList{}

	if err := hl.Load(hostsFile); err != nil {
		return err
	}

	for _, h := range hl.Hosts {
		if _, err := fmt.Fprintln(out, h); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	hostsCmd.AddCommand(listCmd)

	// Aquí definir flags y ajustes de configuración.

	// Cobra admite flags persistentes que funcionarán para
	// el comando y todos los subcomandos, por ejemplo:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra admite flags locales que solo se ejecutarán cuando
	// se llame directamente al comando, por ejemplo:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
