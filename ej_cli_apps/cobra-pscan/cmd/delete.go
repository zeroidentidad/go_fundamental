package cmd

import (
	"fmt"
	"io"
	"os"

	"pscan/scan"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd representa el comando de eliminar
var deleteCmd = &cobra.Command{
	Use:          "delete <host1>...<host n>",
	Aliases:      []string{"d"},
	Short:        "Eliminar hosts de la lista",
	SilenceUsage: true,
	Args:         cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile := viper.GetString("hosts-file")

		return deleteAction(os.Stdout, hostsFile, args)
	},
}

func deleteAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostsList{}

	if err := hl.Load(hostsFile); err != nil {
		return err
	}

	for _, h := range args {
		if err := hl.Remove(h); err != nil {
			return err
		}

		fmt.Fprintln(out, "Deleted host:", h)
	}

	return hl.Save(hostsFile)
}

func init() {
	hostsCmd.AddCommand(deleteCmd)

	// Aquí definir flags y ajustes de configuración.

	// Cobra admite flags persistentes que funcionarán para
	// el comando y todos los subcomandos, por ejemplo:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra admite flags locales que solo se ejecutarán cuando
	// se llame directamente al comando, por ejemplo:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
