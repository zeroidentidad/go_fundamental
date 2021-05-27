package cmd

import (
	"github.com/spf13/cobra"
)

// hostsCmd represents the hosts command
var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Administrar la lista de hosts",
	Long: `Administrar lista de hosts para pScan
	
	Agregar hosts con el comando add
	
	Eliminar hosts con el comando delete
	
	Listar los hosts con el comando list`,
}

func init() {
	rootCmd.AddCommand(hostsCmd)

	// Aquí definir flags y ajustes de configuración.

	// Cobra admite flags persistentes que funcionarán para
	// el comando y todos los subcomandos, por ejemplo:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra admite flags locales que solo se ejecutarán cuando
	// se llame directamente al comando, por ejemplo:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
